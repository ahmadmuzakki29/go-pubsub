package pubsub

import (
	"testing"
	"log"
)

// this test will listen to new message. don't forget to ctrl+c to kill it
var testprojectid = "YOUR GOOGLE CLOUD PROJECT ID HERE"
func TestPublishTopic(t *testing.T) {
	Init(Config{ProjectId:testprojectid})
	
	msg := `{"name":"jeki"}`
	err := Publish("hello", []byte(msg))
	if err != nil {
		t.Error(err)
	}
	
	AddHandler("hello","hellochannel",sendMailHandler)
}

type Person struct{
	Name string `json:"name"`
}

func sendMailHandler(msg *Message)error{
	var p Person
	msg.ToJson(&p)
	log.Println("hello ",p.Name)
	msg.Done(true)
	return nil
}