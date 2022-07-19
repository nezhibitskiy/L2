package main

import (
	"flag"
	"log"

	"10/telnet"
)

var (
	timeout int
	port    string
	host    string
)

//init ...
func init() {
	flag.IntVar(&timeout, "timeout", 15, "timeout for connect")
	flag.StringVar(&port, "port", "", "port for connecting to")
	flag.StringVar(&host, "host", "", "binded ip for connecting to")
}

type Telnet struct {
	client telnet.Telnet
}

func newTelnet(client telnet.Telnet) *Telnet {
	return &Telnet{client: client}
}

func main() {
	flag.Parse()
	host, port = "localhost", "3000"
	myTelnetClient := telnet.NewTelnetClient(host, port, timeout)
	a := newTelnet(myTelnetClient)
	err := a.client.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
