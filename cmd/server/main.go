package main

import (
	"encoding/json"
	"log"

	maelstorm "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/louisandrew/gossip-glomers/internal/messages"
)

func main() {
	n := maelstorm.NewNode()

	n.Handle(messages.Echo, func(msg maelstorm.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// @TODO: add logger
		body["type"] = "echo_ok"
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
