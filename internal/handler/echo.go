package handler

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type echoHandler struct {
	node *maelstrom.Node
}

func (h echoHandler) handle(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	body["type"] = "echo_ok"
	return h.node.Reply(msg, body)
}

func EchoHandler(node *maelstrom.Node) maelstrom.HandlerFunc {
	h := echoHandler{node}
	return h.handle
}
