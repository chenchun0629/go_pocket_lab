package main

import (
	"encoding/json"
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

	//reporter
	producer := xsarama.NewReporter(&cfg, log)
	//subscriber
	//consumer := subscriber.NewSubscriber(&cfg, log)
	message := "Hello Kafka World."

	//ch := make(chan string)
	//consumer.Consume(cfg.Topic, ch)
	producer.DoReport(cfg.Topic, []byte(message))

	//select {
	//case msg := <-ch:
	//	fmt.Println("Got msg: ", msg)
	//	break
	//}
}

