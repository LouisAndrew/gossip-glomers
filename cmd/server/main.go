package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/louisandrew/gossip-glomers/internal/handler"
	"github.com/louisandrew/gossip-glomers/internal/messages"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle(messages.Echo, handler.EchoHandler(n))
	n.Handle(messages.Generate, handler.UniqueIdHandler(n))

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
