package pubsub

import (
	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"log"
)

var publisher = make(map[string]*Publisher)

type Publisher struct {
	topic *pubsub.Topic
}

func getPublisher(topic string) (*Publisher, error) {
	if projectid==""{
		log.Fatal("Init not called!")
	}
	if publisher[topic] == nil {
		ctx := context.Background()

		client, err := pubsub.NewClient(ctx, projectid)
		if err != nil {
			return nil, err
		}

		t := client.Topic(topic)
		if exists, _ := t.Exists(ctx); !exists {
			client.CreateTopic(ctx, topic)
		}

		publisher[topic] = &Publisher{t}
	}
	return publisher[topic], nil
}

func Publish(topic string, msg []byte) error {
	pub, err := getPublisher(topic)
	if err != nil {
		return err
	}

	message := &pubsub.Message{Data: msg}
	_, err = pub.topic.Publish(context.Background(), message)
	return err
}
