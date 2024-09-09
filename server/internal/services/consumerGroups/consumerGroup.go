package consumerGroups

import (
	"fmt"
	"github.com/IBM/sarama"
	m "github.com/titaniper/kafka-admin/internal/models"
	"github.com/titaniper/kafka-admin/pkg/kafka"
	"log"
	"regexp"
)

type Service struct {
	kafkaClient *kafka.KafkaClient
}

func New(kafkaClient *kafka.KafkaClient) *Service {
	return &Service{
		kafkaClient,
	}
}

func (s *Service) List(keyword string, isInactive bool) ([]string, error) {
	// TODO: infrastructure?
	return s.kafkaClient.GetConsumerGroups(keyword, isInactive)
}

func (s *Service) Delete(keyword string) error {
	// TODO: 전체 삭제하지마
	groups, _ := s.List(keyword, true)
	fmt.Printf("Deleting groups: %d\n", len(groups))
	for _, group := range groups {
		fmt.Printf("Deleting group %s\n", group)
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

func (s *Service) GetDetails(groupID string) (*m.ConsumerGroupDetailsResponse, error) {
	// Kafka 클라이언트와 admin 클라이언트 가져오기
	client := s.kafkaClient.Client
	admin := s.kafkaClient.Admin

	// 컨슈머 그룹 설명 가져오기
	groups, err := admin.DescribeConsumerGroups([]string{groupID})
	if err != nil || len(groups) == 0 {
		return nil, fmt.Errorf("failed to describe consumer group: %v", err)
	}
	group := groups[0]

	// 오프셋 정보 가져오기
	offsetFetchResponse, err := admin.ListConsumerGroupOffsets(groupID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list consumer group offsets: %v", err)
	}

	response := &m.ConsumerGroupDetailsResponse{
		Inherit:           "details",
		GroupID:           groupID,
		Members:           len(group.Members),
		Topics:            len(offsetFetchResponse.Blocks),
		Simple:            false,
		PartitionAssignor: group.ProtocolType,
		State:             string(group.State),
		Coordinator: m.Coordinator{
			ID:   int(group.Version),
			Host: group.Protocol,
			Port: int(group.Version),
		},
	}

	var totalLag int64
	for topic, partitions := range offsetFetchResponse.Blocks {
		for partition, block := range partitions {
			endOffset, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
			if err != nil {
				continue
			}

			lag := endOffset - block.Offset
			totalLag += lag

			partitionInfo := m.PartitionInfo{
				Topic:         topic,
				Partition:     partition,
				CurrentOffset: block.Offset,
				EndOffset:     endOffset,
				ConsumerLag:   lag,
			}
			response.Partitions = append(response.Partitions, partitionInfo)
		}
	}

	response.ConsumerLag = totalLag

	// 추가 정보 수집 (예: BytesInPerSec, BytesOutPerSec 등)
	// 이 정보는 Kafka의 JMX 메트릭스나 다른 모니터링 도구를 통해 얻을 수 있습니다.
	// 여기서는 예시로 nil을 설정합니다.
	response.Coordinator.BytesInPerSec = nil
	response.Coordinator.BytesOutPerSec = nil
	response.Coordinator.PartitionsLeader = nil
	response.Coordinator.Partitions = nil
	response.Coordinator.InSyncPartitions = nil
	response.Coordinator.PartitionsSkew = nil
	response.Coordinator.LeadersSkew = nil

	return response, nil
}

func (s *Service) ResetOffset(groupID, topic string, partition int32) error {
	admin := s.kafkaClient.Admin

	err := admin.DeleteConsumerGroupOffset(groupID, topic, partition)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTopicsFromConsumerGroups(consumerGroupPattern, topicPattern string) error {
	admin := s.kafkaClient.Admin

	// 컨슈머 그룹 목록 가져오기
	groups, err := admin.ListConsumerGroups()
	if err != nil {
		return fmt.Errorf("error listing consumer groups: %v", err)
	}

	consumerGroupRegex, err := regexp.Compile(consumerGroupPattern)
	if err != nil {
		return fmt.Errorf("invalid consumer group pattern: %v", err)
	}

	topicRegex, err := regexp.Compile(topicPattern)
	if err != nil {
		return fmt.Errorf("invalid topic pattern: %v", err)
	}

	for group := range groups {
		if consumerGroupRegex.MatchString(group) {
			// 컨슈머 그룹의 오프셋 정보 가져오기
			offsetFetchResponse, err := admin.ListConsumerGroupOffsets(group, nil)
			if err != nil {
				log.Printf("Error listing offsets for group %s: %v", group, err)
				continue
			}

			for topic := range offsetFetchResponse.Blocks {
				if topicRegex.MatchString(topic) {
					// 토픽의 모든 파티션에 대해 오프셋 삭제
					partitions := offsetFetchResponse.Blocks[topic]
					for partition := range partitions {
						err := admin.DeleteConsumerGroupOffset(group, topic, partition)
						if err != nil {
							log.Printf("Error deleting offset for group %s, topic %s, partition %d: %v", group, topic, partition, err)
						} else {
							log.Printf("Successfully deleted offset for group %s, topic %s, partition %d", group, topic, partition)
						}
					}
				}
			}
		}
	}

	return nil
}
