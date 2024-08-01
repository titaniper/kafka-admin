package main

import (
	"fmt"
	"net/http"

	"github.com/titaniper/gopang/libs/kafka"

	cgc "github.com/titaniper/gopang/routers/consumerGroups"
	tc "github.com/titaniper/gopang/routers/topics"
	"github.com/titaniper/gopang/services/consumerGroups"
	"github.com/titaniper/gopang/services/topics"
)

// setJSONContentType는 모든 응답에 Content-Type 헤더를 설정하는 미들웨어입니다.
func setJSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// notFoundHandler는 허용되지 않은 경로에 대한 핸들러입니다.
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 - Not Found")
}

func main() {
	kafkaClient := initKafkaClient()
	startServer(kafkaClient)
}

func initKafkaClient() *kafka.KafkaClient {
	// TODO: 환경 변수
	kafkaClient, err := kafka.New([]string{"kafka-kafka-bootstrap.streaming.svc.cluster.local:9092"})
	if err != nil {
		panic(err)
	}

	return kafkaClient
}

// TODO: router interface만 받도록 수정
func startServer(kafkaClient *kafka.KafkaClient) {
	// TODO: DI
	topicService := topics.New(kafkaClient) // 서비스 초기화 예제
	consumerGroupsService := consumerGroups.New(kafkaClient)
	topicController := tc.New(topicService)
	consumerGroupsController := cgc.New(consumerGroupsService)

	// New Server Multiplexer
	// 네트워크와 전자공학에서 멀티플렉서는 여러 신호를 하나의 신호로 결합하는 장치입니다. 반대로 디멀티플렉서(Demultiplexer)는 하나의 신호를 여러 신호로 분리합니다.
	mux := http.NewServeMux()

	// 라우트 설정
	for _, route := range topicController.Routes() {
		mux.HandleFunc(route.Path, route.Handler)
	}
	// TODO: 배열 하나에 담자?
	for _, route := range consumerGroupsController.Routes() {
		mux.HandleFunc(route.Path, route.Handler)
	}

	// 기본 핸들러 외의 모든 경로에 대해 notFoundHandler 설정
	mux.HandleFunc("/", notFoundHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// 미들웨어를 사용하여 Content-Type 설정
	wrappedMux := setJSONContentType(mux)

	fmt.Println("Starting server at port 8080")
	// TODO: 환경 변수
	if err := http.ListenAndServe(":8080", wrappedMux); err != nil {
		fmt.Println(err)
	}
}
