package kafka

import (
	"github.com/IBM/sarama"
	"log"
)

type KafkaClient struct {
	Client sarama.Client
	Admin  sarama.ClusterAdmin
}

func New(borkers []string) (*KafkaClient, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_7_2_0 // Kafka 버전에 맞춰 설정

	// Kafka 클라이언트 생성
	client, err := sarama.NewClient(borkers, config)
	if err != nil {
		log.Fatalf("Error creating Kafka Client: %v", err)
	}
	//defer Client.Close()

	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		log.Fatalf("Error creating Kafka cluster Admin: %v", err)
	}
	//defer Admin.Close()

	return &KafkaClient{
		client,
		admin,
	}, nil
}
