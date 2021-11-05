package xsarama

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

type Subscriber struct {
	Consumer sarama.Consumer
	logger   *logrus.Logger
}

func NewSubscriber(cfg *KafkaCfg, log *logrus.Logger) *Subscriber {
	subscriber := &Subscriber{
		logger: log,
	}
	subscriber.setConsumer(cfg)
	return subscriber
}

func (subscriber *Subscriber) setConsumer(cfg *KafkaCfg) {
	consumer, err := sarama.NewConsumer([]string{cfg.Host}, nil)
	if err != nil {
		panic(err)
	}
	subscriber.Consumer = consumer
}

func (subscriber *Subscriber) Consume(topic string, ch chan string) {
	defer func() {
		if err := subscriber.Consumer.Close(); err != nil {
			subscriber.logger.Errorf("SubcriberCloseConsumerError", err)
		}
	}()

	// 获取所有 partition
	partitionList, err := subscriber.Consumer.Partitions(topic)
	if err != nil {
		subscriber.logger.Errorf("SubScriberGetPartitionsError", err, topic)
	}

	// 遍历所有 partition 获取最新 offset 上的消息
	for _, partition := range partitionList {
		pc, _ := subscriber.Consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		// 启动一个 goroutine 用于持续监听消息
		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				subscriber.logger.Infof(messageReceived(message))
				ch <- messageReceived(message)
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) string {
	return string(message.Value)
}
