package main

import (
	"flag"
	"log"

	"github.com/sahlinet/go-tumbo/pkg/app"
	"github.com/sahlinet/go-tumbo/pkg/client"
	"github.com/sahlinet/go-tumbo/pkg/config"
)

func main() {
	var server = flag.Bool("server", false, "huhu")
	flag.Parse()
	log.Printf("Running as server: %t", *server)

	/*if *server {
		srv.Start()
	}
	*/

	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run("0.0.0.0:3000")

	client.GetWorker()
}
