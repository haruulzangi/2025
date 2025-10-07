package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/haruulzangi/2025/challenges/round-2/crypto/lost-file/challenge/flag"
	"golang.org/x/crypto/ssh"
)

func ListenAndServe(flagTemplate string, listenAddr string, oracleAddress string) error {
	config, err := prepareServerConfig(oracleAddress)
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
			go ssh.DiscardRequests(reqs)
			for newChannel := range chans {
				if newChannel.ChannelType() != "session" {
					newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
					continue
				}
				channel, requests, err := newChannel.Accept()
				if err != nil {
					continue
				}
				go ssh.DiscardRequests(requests)

				message := fmt.Sprintln(flag.GetFlag(flagTemplate, sshConn))
				io.WriteString(channel, message)
				channel.Close()
			}
		}(nConn)
	}
}
