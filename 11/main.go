package main

import (
	"log"

	"11/internal"
	"11/internal/hash"
	"11/internal/server"
)

func main() {
	config := internal.NewConfig()

	hash, err := hash.NewHash()
	if err != nil {
		log.Fatal(err)
	}

	app := server.NewApp(config, hash)

	if err = app.Start(); err != nil {
		log.Fatal(err)
	}
}
