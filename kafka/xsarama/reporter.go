package xsarama

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// 消息生产者数据结构
type Reporter struct {
	Producer sarama.SyncProducer
	logger   *logrus.Logger
}

// 构造函数
func NewReporter(cfg *KafkaCfg, log *logrus.Logger) *Reporter {
	reporter := &Reporter{
		logger: log,
	}

	reporter.setProducer(cfg)
	return reporter
}

func (reporter *Reporter) setProducer(cfg *KafkaCfg) {
	var broker = []string{cfg.Host}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	// new 一个Producer
	producer, err := sarama.NewSyncProducer(broker, config)
	if err != nil {
		reporter.logger.Errorf("new producer error", err)
	}

	reporter.Producer = producer
}

func (reporter *Reporter) DoReport(topic string, msg []byte) {
	reporter.do(topic, msg)
}

func (reporter *Reporter) do(topic string, msg []byte) {
	kafkaMsg := generateProducerMessage(topic, msg)
	_, _, err := reporter.Producer.SendMessage(kafkaMsg)
	if err != nil {
		reporter.logger.Errorf("producer send message error", err, string(msg))
	}
	reporter.logger.Infof("producer send message", string(msg))
}

func generateProducerMessage(topic string, message []byte) *sarama.ProducerMessage {
	return &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}
}

