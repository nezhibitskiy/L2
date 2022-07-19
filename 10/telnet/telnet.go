package telnet

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

type TelnetClient struct {
	conn    net.Conn
	host    string
	port    string
	timeout time.Duration
}

type Telnet interface {
	Connect() error
}

func NewTelnetClient(host, port string, timeout int) *TelnetClient {
	return &TelnetClient{
		host:    host,
		port:    port,
		timeout: time.Duration(timeout) * time.Second,
	}
}

//Connect ...
func (t *TelnetClient) Connect() (err error) {
	t.conn, err = net.DialTimeout("tcp", net.JoinHostPort(t.host, t.port), t.timeout)
	if err != nil {
		return err
	}

	defer func() {
		if err := t.conn.Close(); err != nil {
			panic(err)
		}
	}()

	return t.goClientWriter()
}

func (t *TelnetClient) goClientWriter() (err error) {
	errChan := make(chan error, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	read := func() {
		for {
			select {
			case <-ctx.Done():
				return

			default:
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					input := scanner.Bytes()
					err = t.conn.SetDeadline(time.Now().Add(t.timeout))
					if err != nil {
						errChan <- err
					}

					_, err = t.conn.Write(input)
					if err != nil {
						errChan <- err
					}

					out := []byte{}

					data, err := t.conn.Read(out)
					if err != nil {
						errChan <- err
					}

					fmt.Println(data)
				}
			}
		}
	}

	go read()

	for {
		select {
		case x := <-errChan:
			ctx.Done()
			return x
		default:
			time.Sleep(10 * time.Second)
		}
	}

}
