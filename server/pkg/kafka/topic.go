package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"strings"
)

func (c *KafkaClient) CreateTopic(name string) error {
	return c.Admin.CreateTopic(name, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
}

func (c *KafkaClient) GetTopics(keyword string) ([]string, error) {
	topics, err := c.Admin.ListTopics()
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

func (c *KafkaClient) DeleteTopic(name string) error {
	return c.Admin.DeleteTopic(name)
}
