package handler

import (
	"encoding/json"
	"fmt"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type maelstromResponse map[string]any
type maelstromHandler func(msg maelstrom.Message) (maelstromResponse, error)

func getMessageBody(msg maelstrom.Message, expectedType string) (maelstromResponse, error) {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return nil, err
	}

	if body["type"] != expectedType {
		return nil, fmt.Errorf("unexpected message type. Expected: %s. Got: %s", expectedType, body["type"])
	}

	return body, nil
}

func buildHandler(node *maelstrom.Node, handler maelstromHandler) maelstrom.HandlerFunc {
	return func(msg maelstrom.Message) error {
		response, err := handler(msg)
		if err != nil {
			return err
		}

		return node.Reply(msg, response)
	}
}

func buildTestMessage(data maelstromResponse) maelstrom.Message {
	body, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return maelstrom.Message{
		Src:  "src",
		Dest: "dest",
		Body: body,
	}
}
