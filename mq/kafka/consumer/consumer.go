package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "mygroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	c.SubscribeTopics([]string{"first-topic", "^aRegex.*[Tt]opic"}, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer err: %v (%v)\n", err, msg)
		}
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}
	c.Close()
}
