package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	topic := "123"
	ip := "localhost:9092"

	consumer, err := sarama.NewConsumer([]string{ip}, nil)
	if err != nil {
		fmt.Printf("Failed to start consumer: %s\n", err)
		return
	}
	defer consumer.Close()

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("failed to get the list of partitions:", err)
		return
	}
	fmt.Println("partitionList:", partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		fmt.Printf("consumer consume partition: %d\n", partition)
		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			defer wg.Done()
			defer pc.AsyncClose()
			for msg := range pc.Messages() {
				fmt.Printf("msg: topic: %s, offset: %d, key: %s, value: %s\n", msg.Topic, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}

	wg.Wait()
}
