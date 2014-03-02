package sse

import (
	"fmt"
	"net/http"
)

func NewBroker() *Broker {
	return &Broker{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}
}

type Broker struct {
	clients        map[chan string]bool
	newClients     chan chan string
	defunctClients chan chan string
	Messages       chan string
}

func (broker *Broker) Start() {
	for {
		select {
		case client := <-broker.newClients:
			broker.clients[client] = true
		case client := <-broker.defunctClients:
			delete(broker.clients, client)
		case message := <-broker.Messages:
			for client, _ := range broker.clients {
				client <- message
			}
		}
	}
}

func (broker *Broker) Handler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		flusher, ok := writer.(http.Flusher) // ensure supports streaming
		if !ok {
			http.Error(writer, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		client := make(chan string)
		broker.newClients <- client

		defer func() {
			broker.defunctClients <- client
		}()

		writer.Header().Set("Content-Type", "text/event-stream")
		writer.Header().Set("Cache-Control", "no-cache")
		writer.Header().Set("Connection", "keep-alive")

		// Don't close the connection, instead loop 10 times,
		// sending messages and flushing the response each time
		// there is a new message to send along.
		//
		// NOTE: we could loop endlessly; however, then you
		// could not easily detect clients that dettach and the
		// server would continue to send them messages long after
		// they're gone due to the "keep-alive" header.  One of
		// the nifty aspects of SSE is that clients automatically
		// reconnect when they lose their connection.
		//
		// A better way to do this is to use the CloseNotifier
		// interface that will appear in future releases of
		// Go (this is written as of 1.0.3):
		// https://code.google.com/p/go/source/detail?name=3292433291b2
		//
		for i := 0; i < 10; i++ {
			msg := <-client
			fmt.Fprintf(writer, "data: %s\n\n", msg)
			flusher.Flush()
		}
	}
}
