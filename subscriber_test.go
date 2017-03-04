package pubsub

import (
	"fmt"
	"testing"
)

func TestPublishTopic(t *testing.T) {
	msg := "Hello World"
	err := Publish("send_mail", []byte(msg))
	fmt.Println(err)
}
