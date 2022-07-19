package main

import (
	"flag"
	"log"

	"11/internal/config"
	"11/internal/hash"
	"11/internal/server"
)

func main() {
	flag.Parse()

	conf := config.NewConfig()

	hash, err := hash.NewHash()
	if err != nil {
		log.Fatal(err)
	}

	app := server.NewApp(conf, hash)

	if err = app.Start(); err != nil {
		log.Fatal(err)
	}
}

func catchSig() {

}
