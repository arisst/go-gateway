package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	message := Message(true, "test message")

	assert.Equal(t, message["status"], true)
	assert.Equal(t, message["message"], "test message")
}

func TestRequest(t *testing.T) {
	request := Request("https://run.mocky.io/v3/24524fbf-3d9b-47d0-bd47-060c66e5645f")

	assert.Equal(t, request["status"], true)
}
