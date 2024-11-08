package handler

import (
	"reflect"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	got, err := handleEcho(buildTestMessage(maelstromResponse{"type": "echo"}))
	want := maelstromResponse{"type": "echo_ok"}

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
