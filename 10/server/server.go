package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	port = "3000"
	host = "localhost"
)

func main() {
	srv, err := net.Listen("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()

	for {
		sigConn, err := srv.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func(sigConn net.Conn) {
			for {
				data, err := bufio.NewReader(sigConn).ReadString('\n')
				if err != nil {
					fmt.Println(err)
				}

				_, err = sigConn.Write([]byte(fmt.Sprintf("%s year now: %s", data, time.Now().Format("2006"))))

			}
		}(sigConn)
	}
}
