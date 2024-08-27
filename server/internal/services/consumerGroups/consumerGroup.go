package consumerGroups

import (
	"github.com/titaniper/kafka-admin/pkg/kafka"
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

	// NOTE: haulla-api--internal-stage

	/**
	1. 컨슈머 그룹 가져옴
	- haulla-api-internal-stage 제외
	- haulla-api-5285919834-internal-stage 패턴들
	*/

	/**
	  2. 컨슈머 그룹에서
	*/
	// partitioned.haulla-5285919834.domain_event

	// TODO: infrastructure?
	//return s.kafkaClient.DeleteConsumerGroup(keyword)
	return nil
}
