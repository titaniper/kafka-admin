package kafka

import (
	"log"
	"strings"
)

func (c *KafkaClient) GetConsumerGroups(keyword string, isInactive bool) ([]string, error) {
	consumerGroups, err := c.Admin.ListConsumerGroups()
	if err != nil {
		log.Fatalf("Error listing consumerGroups: %v", err)
		return nil, err
	}

	var filtered []string
	for consumerGroup := range consumerGroups {
		if keyword == "" || strings.Contains(consumerGroup, keyword) {
			if isInactive {
				// Describe the consumer group to get its members
				descriptions, err := c.Admin.DescribeConsumerGroups([]string{consumerGroup})
				if err != nil {
					log.Printf("Error describing consumer group %s: %v", consumerGroup, err)
					continue
				}

				// Check if the consumer group has any members
				if len(descriptions[0].Members) > 0 {
					continue // Skip this group as it has active members
				}
			}

			filtered = append(filtered, consumerGroup)
		}
	}
	return filtered, nil
}

func (c *KafkaClient) DeleteConsumerGroup(name string) error {
	return c.Admin.DeleteConsumerGroup(name)
}

func (c *KafkaClient) DeleteConsumerGroupOffset(name, topic string, partition int32) error {
	return c.Admin.DeleteConsumerGroupOffset(name, topic, partition)
}
