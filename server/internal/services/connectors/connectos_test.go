package connectors

import (
	"fmt"
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
	client := New("http://localhost:56960/")

	response, _ := client.GetAllConnector()
	for _, name := range response.Connectors {
		connector, _ := client.GetConnector(name)

		statusResponse, _ := client.GetConnectorStatus(name)
		if len(connector.Tasks) > 0 {
			taskStatus := statusResponse.TasksStatus[0]
			if taskStatus.State == "FAILED" {
				println("connector.Name", connector.Name)
				errorType := 0
				if strings.Contains(taskStatus.Trace, " An exception occurred in the change event producer. This connector will be stopped.") {
					errorType = 1
				} else if strings.Contains(taskStatus.Trace, "The database schema history couldn't be recovered. Consider to increase the value for schema.history.internal.kafka.recovery.poll.interval.ms") {
					errorType = 2
				} else if strings.Contains(taskStatus.Trace, "java.lang.OutOfMemoryError: Java heap space") {
					errorType = 3
				}

				if errorType > 0 {
					fmt.Printf("type %d", errorType)
					println("trace", taskStatus.Trace)
					client.RestartTask(connector.Name, taskStatus.ID)
				}
			}
		}
	}
}
