package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/go-nats"
)

func main() {
	natsURI := os.Getenv("NATS_URI")
	hostname, _ := os.Hostname()
	log.Printf("NATS address: %s, hostname: %s", natsURI, hostname)

	nc, err := nats.Connect(natsURI)
	if err != nil {
		log.Fatalf("Connect to NATS failed: %v", err)
	}

	log.Println("Connected to NATS at:", nc.ConnectedUrl())

	nc.Subscribe("tasks", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte(fmt.Sprintf("Hello from host %s", hostname)))
	})

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Kill, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	log.Println("Api server will shutdown...")
}
