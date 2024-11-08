package handler

import (
	"crypto/rand"
	"encoding/base32"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/louisandrew/gossip-glomers/internal/messages"
)

// make it longer for a larger amount of possible unique ids
const idLength = 16

func generateUniqueId() (string, error) {
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	// encode randomBytes to a string
	return base32.StdEncoding.EncodeToString(randomBytes)[:idLength], nil
}

func handleUniqueId(msg maelstrom.Message) (maelstromResponse, error) {
	body, err := getMessageBody(msg, messages.Generate)
	if err != nil {
		return nil, err
	}

	uniqueId, err := generateUniqueId()
	if err != nil {
		return nil, err
	}

	body["id"] = uniqueId
	body["type"] = "generate_ok"
	return body, nil
}

func UniqueIdHandler(node *maelstrom.Node) maelstrom.HandlerFunc {
	return buildHandler(node, handleUniqueId)
}
