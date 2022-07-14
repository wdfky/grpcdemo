package main

import (
	"flag"
	"log"

	"github.com/wdfky/grpcdemo/client"
	"github.com/wdfky/grpcdemo/server"
)

var mode = flag.String("mode", "server", "server or client")

func main() {
	flag.Parse()
	switch *mode {
	case "server":
		server.Init()
	case "client":
		client.Init()
	default:
		log.Fatal("no match mod")
	}
}
