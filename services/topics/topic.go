package topics

import (
	"github.com/titaniper/gopang/libs/kafka"
)

type Service struct {
	kafkaClient *kafka.KafkaClient
}

func New(kafkaClient *kafka.KafkaClient) *Service {
	return &Service{
		kafkaClient,
	}
}

func (s *Service) CreateTopic(name string) error {
	return s.kafkaClient.CreateTopic(name)
}

func (s *Service) GetTopics(keyword string) ([]string, error) {
	// TODO: infrastructure?
	return s.kafkaClient.GetTopics(keyword)
}
