package kafka

import (
	"github.com/megoo/mq/kafka/consumer"
	"github.com/megoo/mq/kafka/producer"
	"testing"
)

func TestKafka(t *testing.T) {
	consumer.Consumer()
	producer.Producer()
}
