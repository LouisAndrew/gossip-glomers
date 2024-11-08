package handler

import (
	"testing"
)

const caseNum = 100

var ids = make(map[string]bool)

func TestUniqueIdHandler(t *testing.T) {
	for i := 0; i < caseNum; i++ {
		got, err := handleUniqueId(buildTestMessage(maelstromResponse{"type": "generate"}))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if got["type"] != "generate_ok" {
			t.Errorf("unexpected response: %v", got)
		}

		generatedId := got["id"].(string)
		if _, ok := ids[generatedId]; ok {
			t.Errorf("id %s is not unique", generatedId)
		} else {
			ids[generatedId] = true
		}
	}
}
