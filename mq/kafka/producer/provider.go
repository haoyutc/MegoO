package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producer() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	go func() {
		for event := range p.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				}
				fmt.Printf("Delivery message to %v\n", ev.TopicPartition)
			}
		}
	}()
	topic := "first-topic"
	for _, word := range []string{"welcome ", "to", "the", "confluent", "kafka", "golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(word),
		}, nil)
	}
	p.Flush(1000 * 15)
}
