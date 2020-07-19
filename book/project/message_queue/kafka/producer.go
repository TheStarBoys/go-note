package main

import (
	"github.com/Shopify/sarama"
	"time"
	"fmt"
)

func main() {
	topic := "123"
	ip := "localhost:9092"
	config := sarama.NewConfig()
	// 设置 kafka 回不回ack，如果不回，可能会被传丢
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 分区：
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 这个消息就是要写入 kafka 的消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic

	// 同步的客户端
	client, err := sarama.NewSyncProducer([]string{ip}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	count := 1
	for {
		msg.Value = sarama.StringEncoder(fmt.Sprintf("this is a good test, my message is good. count=%d", count))
		count++
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed, err:", err)
			return
		}
		fmt.Printf("partition_id: %d, offset: %d\n", pid, offset)
		time.Sleep(1 * time.Second)
	}
}
