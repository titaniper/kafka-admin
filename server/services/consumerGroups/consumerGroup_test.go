package consumerGroups

import (
	"github.com/titaniper/kafka-admin/libs/kafka"
	"strings"
	"testing"
)

// 열거형 정의
type ConnectorError int

const (
	Metadata ConnectorError = iota
)

// TODO: 개선 ㅋㅋ
func Test_GET_CONNECTOR(t *testing.T) {
	kafkaClient, _ := kafka.New([]string{"kafka-kafka-bootstrap.streaming.svc.cluster.local:9092"})
	client := New(kafkaClient)

	// 1.
	response, _ := client.List()
	for _, name := range response.Connectors {
		connector, _ := client.GetConnector(name)

		statusResponse, _ := client.GetConnectorStatus(name)
		if len(connector.Tasks) > 0 {
			taskStatus := statusResponse.TasksStatus[0]
			if taskStatus.State == "FAILED" {
				if strings.Contains(taskStatus.Trace, " An exception occurred in the change event producer. This connector will be stopped.") {
					println("Metadata error", connector.Name)
					client.RestartTask(connector.Name, taskStatus.ID)
				} else {
					println("Another", connector.Name)
				}
			}
		}
	}
}