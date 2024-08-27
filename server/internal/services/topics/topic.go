package topics

import (
	"fmt"
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

func (s *Service) CreateTopic(name string) error {
	return s.kafkaClient.CreateTopic(name)
}

func (s *Service) GetTopics(keyword string) ([]string, error) {
	// TODO: infrastructure?
	return s.kafkaClient.GetTopics(keyword)
}

func (s *Service) GetConsumerGroupTopics(groupID string) ([]string, error) {
	// 컨슈머 그룹의 오프셋 정보 가져오기
	offsetFetchResponse, err := s.kafkaClient.Admin.ListConsumerGroupOffsets(groupID, nil)
	if err != nil {
		return nil, fmt.Errorf("error listing consumer group offsets: %v", err)
	}

	// 구독 중인 토픽 목록 추출
	subscribedTopics := make(map[string]struct{})
	for topic, partitionOffsets := range offsetFetchResponse.Blocks {
		// 파티션 오프셋이 하나라도 있으면 해당 토픽을 구독 중인 것으로 간주
		if len(partitionOffsets) > 0 {
			subscribedTopics[topic] = struct{}{}
		}
	}

	// 결과 슬라이스 생성
	result := make([]string, 0, len(subscribedTopics))
	for topic := range subscribedTopics {
		result = append(result, topic)
	}

	return result, nil
}
