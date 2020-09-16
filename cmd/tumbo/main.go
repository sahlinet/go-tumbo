package main

import (
	"flag"
	"log"

	"github.com/sahlinet/tumbo/pkg/client"
	srv "github.com/sahlinet/tumbo/pkg/server"
)

func main() {
	var server = flag.Bool("server", false, "huhu")
	flag.Parse()
	log.Printf("Running as server: %t", *server)

	if *server {
		srv.Start()
	}

	client.GetWorker()
}
