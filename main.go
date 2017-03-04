package pubsub

import (
	"log"
	"cloud.google.com/go/pubsub"
	"encoding/json"
)

type Config struct{
	ProjectId string
	Output func(...interface{})
}

var config *Config

func Init(cfg Config){
	projectid = cfg.ProjectId
	if cfg.Output==nil{
		cfg.Output = output
	}
	config = &cfg
}

func output(msg ...interface{}){
	log.Println(msg...)
}

type Message struct {
	pubsub.Message
}

func (m Message) ToJson(dest interface{})error{
	return json.Unmarshal(m.Data,dest)
}