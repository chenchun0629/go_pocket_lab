package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"pocket_lab/kafka/xsarama"
)

func main() {
	var (
		config = `{
	"host": "localhost:9092",
	"topic": "kafka_test"
  }`
	)

	var cfg xsarama.KafkaCfg
	_ = json.Unmarshal([]byte(config), &cfg)
	log := &logrus.Logger{Out: os.Stdout}

	//subscriber
	consumer := xsarama.NewSubscriber(&cfg, log)

	ch := make(chan string)
	consumer.Consume(cfg.Topic, ch)

	select {
	case msg := <-ch:
		fmt.Println("Got msg: ", msg)
		break
	}
}

