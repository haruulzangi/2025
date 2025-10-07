package flag

import (
	"fmt"
	"log"
	"net"
	"time"

	cfg "github.com/haruulzangi/2025/challenges/round-3/misc/knock-knock/challenge/internal/cfg"
	"golang.org/x/crypto/ssh"
)

func acceptConnectionOrDispose(flag string, listener net.Listener) {
	defer listener.Close()

	connChan := make(chan net.Conn)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			connChan <- conn
		}
	}()

	select {
	case <-time.After(cfg.CONNECTION_TIMEOUT):
		return
	case conn := <-connChan:
		log.Printf("Flag was accessed: %s", flag)
		fmt.Fprintf(conn, "Nice one, here is the flag: %s\n", flag)
		conn.Close()
	}
}

func SpawnFlagServer(ctx ssh.ConnMetadata) (int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	go acceptConnectionOrDispose(GetFlag(ctx), listener)
	return listener.Addr().(*net.TCPAddr).Port, nil
}
