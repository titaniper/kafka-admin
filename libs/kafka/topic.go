package kafka

import (
	"github.com/IBM/sarama"
	"log"
	"strings"
)

func (c *KafkaClient) CreateTopic(name string) error {
	return c.admin.CreateTopic(name, &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
}

func (c *KafkaClient) GetTopics(keyword string) ([]string, error) {
	topics, err := c.admin.ListTopics()
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
