package kafka

import (
	"github.com/IBM/sarama"
	"log"
)

type KafkaClient struct {
	client sarama.Client
	admin  sarama.ClusterAdmin
}

func New(borkers []string) (*KafkaClient, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0 // Kafka 버전에 맞춰 설정

	// Kafka 클라이언트 생성
	client, err := sarama.NewClient(borkers, config)
	if err != nil {
		log.Fatalf("Error creating Kafka client: %v", err)
	}
	//defer client.Close()

	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		log.Fatalf("Error creating Kafka cluster admin: %v", err)
	}
	//defer admin.Close()

	return &KafkaClient{
		client,
		admin,
	}, nil
}
