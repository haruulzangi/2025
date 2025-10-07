package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	cfg "github.com/haruulzangi/2025/challenges/round-3/misc/knock-knock/challenge/internal/cfg"
	"golang.org/x/crypto/ssh"
)

func ListenAndServe(listenAddr string) error {
	config, err := prepareServerConfig()
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %s", listenAddr, err)
	}
	log.Printf("Listening on %s...", listenAddr)

	for {
		nConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept incoming connection: %s", err)
			continue
		}
		go func(nConn net.Conn) {
			sshConn, chans, reqs, err := ssh.NewServerConn(nConn, config)
			if err != nil {
				nConn.Close()
				return
			}

			defer sshConn.Close()
			go func() {
				<-time.After(cfg.CONNECTION_TIMEOUT)
				sshConn.Close()
			}()

			go ssh.DiscardRequests(reqs)
			for ch := range chans {
				switch ch.ChannelType() {
				case "direct-tcpip":
					channel, requests, err := ch.Accept()
					if err != nil {
						log.Printf("direct-tcpip accept failed: %v", err)
						continue
					}

					var payload struct {
						DestAddr string
						DestPort uint32
						OrigAddr string
						OrigPort uint32
					}
					ssh.Unmarshal(ch.ExtraData(), &payload)

					target := net.JoinHostPort(payload.DestAddr, fmt.Sprint(payload.DestPort))
					log.Printf("forwarding request to %s", target)

					dst, err := net.Dial("tcp", target)
					if err != nil {
						log.Printf("failed to connect to %s: %v", target, err)
						channel.Close()
						continue
					}

					// forward data both ways
					go io.Copy(channel, dst)
					go io.Copy(dst, channel)

					// discard channel-specific requests
					go ssh.DiscardRequests(requests)
				default:
					ch.Reject(ssh.UnknownChannelType, "Prohibited")
				}
			}
		}(nConn)
	}
}
