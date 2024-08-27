package connectors

import (
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
	client := New("http://localhost:51541")

	response, _ := client.GetAllConnector()
	for _, name := range response.Connectors {
		connector, _ := client.GetConnector(name)

		statusResponse, _ := client.GetConnectorStatus(name)
		if len(connector.Tasks) > 0 {
			taskStatus := statusResponse.TasksStatus[0]
			if taskStatus.State == "FAILED" {
				if strings.Contains(taskStatus.Trace, " An exception occurred in the change event producer. This connector will be stopped.") {
					println("1 Metadata error", connector.Name)
					client.RestartTask(connector.Name, taskStatus.ID)
				} else if strings.Contains(taskStatus.Trace, "The database schema history couldn't be recovered. Consider to increase the value for schema.history.internal.kafka.recovery.poll.interval.ms") {
					println("2 Metadata error", connector.Name)
					client.RestartTask(connector.Name, taskStatus.ID)
				} else {
					println("Another", connector.Name)
				}
			}
		}
	}
}
