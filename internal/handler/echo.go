package handler

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/louisandrew/gossip-glomers/internal/messages"
)

func handleEcho(msg maelstrom.Message) (maelstromResponse, error) {
	body, err := getMessageBody(msg, messages.Echo)
	if err != nil {
		return nil, err
	}

	body["type"] = "echo_ok"
	return body, nil
}

func EchoHandler(node *maelstrom.Node) maelstrom.HandlerFunc {
	return buildHandler(node, handleEcho)
}
