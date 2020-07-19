package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) error {
	var err error
	config := sarama.NewConfig()
	// 设置 kafka 回不回ack，如果不回，可能会被传丢
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 分区：
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 同步的客户端
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Warn("init kafka producer failed, err:", err)
		return err
	}

	logs.Debug("init kafka success")
	return nil
}

func Send(data, topic string) error {
	// 这个消息就是要写入 kafka 的消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send message failed, err:", err)
		return err
	}
	logs.Debug("send to kafka succ, partition_id: %d, offset: %d, topic: %s", pid, offset, topic)
	return nil
}