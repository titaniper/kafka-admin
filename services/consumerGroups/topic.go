package consumerGroups

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

func (s *Service) List(keyword string) ([]string, error) {
	// TODO: infrastructure?
	return s.kafkaClient.GetConsumerGroups(keyword)
}

func (s *Service) Delete(keyword string) error {
	// TODO: 전체 삭제하지마
	groups, _ := s.List(keyword)
	for _, group := range groups {
		s.kafkaClient.DeleteConsumerGroup(group)
	}

	// TODO: infrastructure?
	//return s.kafkaClient.DeleteConsumerGroup(keyword)
	return nil
}
