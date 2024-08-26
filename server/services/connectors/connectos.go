package connectors

import (
	connectors2 "github.com/ricardo-ch/go-kafka-connect/lib/connectors"
)

type Service struct {
	kafkaConnectClient connectors2.HighLevelClient
}

func New(url string) *Service {
	return &Service{
		connectors2.NewClient(url),
	}
}

func (s *Service) GetAllConnector() (connectors2.GetAllConnectorsResponse, error) {
	return s.kafkaConnectClient.GetAll()
}

func (s *Service) GetConnector(name string) (connectors2.ConnectorResponse, error) {
	return s.kafkaConnectClient.GetConnector(connectors2.ConnectorRequest{
		name,
	})
}

func (s *Service) GetConnectorStatus(name string) (connectors2.GetConnectorStatusResponse, error) {
	return s.kafkaConnectClient.GetConnectorStatus(connectors2.ConnectorRequest{
		name,
	})
}

func (s *Service) GetTaskStatus(Connector string, TaskID int) (connectors2.TaskStatusResponse, error) {
	return s.kafkaConnectClient.GetTaskStatus(connectors2.TaskRequest{
		Connector,
		TaskID,
	})
}

func (s *Service) RestartTask(Connector string, TaskID int) (connectors2.EmptyResponse, error) {
	return s.kafkaConnectClient.RestartTask(connectors2.TaskRequest{
		Connector,
		TaskID,
	})
}
