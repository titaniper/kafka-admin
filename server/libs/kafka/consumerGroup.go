package kafka

import (
	"log"
	"strings"
)

func (c *KafkaClient) GetConsumerGroups(keyword string) ([]string, error) {
	topics, err := c.admin.ListConsumerGroups()
	if err != nil {
		log.Fatalf("Error listing topics: %v", err)
		return nil, err
	}

	var filtered []string
	for topic := range topics {
		if keyword == "" || strings.Contains(topic, keyword) {
			filtered = append(filtered, topic)
		}
	}
	return filtered, nil
}

func (c *KafkaClient) DeleteConsumerGroup(name string) error {
	return c.admin.DeleteConsumerGroup(name)
}

func (c *KafkaClient) DeleteConsumerGroupOffset(name, topic string, partition int32) error {
	return c.admin.DeleteConsumerGroupOffset(name, topic, partition)
}
