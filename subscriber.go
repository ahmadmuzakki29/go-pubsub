package pubsub

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var client *pubsub.Client

var projectid string

func AddHandler(topic, channel string, handler HandlerFunc) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	var err error
	client, err = pubsub.NewClient(ctx, projectid)
	if err != nil {
		log.Fatal(err)
	}

	t := client.Topic(topic)
	subscriber := client.Subscription(channel)
	if exists, err := subscriber.Exists(ctx); !exists || err != nil {
		fmt.Println("topic:"+topic+" channel:"+channel+" doesn't exists. creating one. err:", err)
		subscriber, err = client.CreateSubscription(ctx, channel, t, 30*time.Second, nil)

		if err != nil {
			log.Fatal(err)
		}
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go subscribe(ctx, subscriber, handler)

	<-quit
	fmt.Println("Exiting subscriber")
	cancel()
	client.Close()
}

type HandlerFunc func(message *pubsub.Message) error

func subscribe(ctx context.Context, subscriber *pubsub.Subscription, handler HandlerFunc) {
	iterator, err := subscriber.Pull(ctx)

	topic := getTopic(ctx, subscriber)
	if err != nil {
		log.Println(subscriber.String(), "/", topic, " ERROR ", err)
		return
	}

	fmt.Println("subscribing topic :", topic)

	for {
		message, err := iterator.Next()
		if err != nil {
			fmt.Println(err)
			break
		}

		err = handler(message)
		if err != nil {
			config.Output(subscriber.String(), "/", topic, " ERROR ", err)
		}
	}
}

func getTopic(ctx context.Context, sub *pubsub.Subscription) string {
	conf, err := sub.Config(ctx)
	if err != nil {
		log.Println(err)
		return ""
	}
	return conf.Topic.String()
}
