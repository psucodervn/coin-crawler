package main

import (
  "github.com/nats-io/go-nats"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"
)

func main() {
	natsURI := os.Getenv("NATS_URI")
	log.Println("NATS address:", natsURI)

	nc, err := nats.Connect(natsURI)
	if err != nil {
		log.Fatalf("Connect to NATS failed: %v", err)
	}

	log.Println("Connected to NATS at:", nc.ConnectedUrl())

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
    resp, err := nc.Request("tasks", []byte("help please"), 10*time.Second)
    if err != nil {
      log.Fatalf("Request failed: %v", err)
    }
    data := string(resp.Data)
    log.Printf("Response: %v", data)
    writer.Write(resp.Data)
  })

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Kill, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	log.Println("Api server will shutdown...")
}
