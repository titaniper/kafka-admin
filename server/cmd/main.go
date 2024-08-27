package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/titaniper/kafka-admin/docs" // Swagger 문서를 임포트합니다
	cgc "github.com/titaniper/kafka-admin/internal/routers/consumerGroups"
	tc "github.com/titaniper/kafka-admin/internal/routers/topics"
	"github.com/titaniper/kafka-admin/internal/services/consumerGroups"
	"github.com/titaniper/kafka-admin/internal/services/topics"
	"github.com/titaniper/kafka-admin/pkg/kafka"
)

// @title Kafka Admin API
// @version 1.0
// @description This is a Kafka admin service API
// @host localhost:8080
// @BasePath /
func main() {
	kafkaClient := initKafkaClient()
	startServer(kafkaClient)
}

func initKafkaClient() *kafka.KafkaClient {
	kafkaClient, err := kafka.New([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}
	return kafkaClient
}

func startServer(kafkaClient *kafka.KafkaClient) {
	r := gin.Default()

	// Swagger 설정
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	topicService := topics.New(kafkaClient)
	consumerGroupsService := consumerGroups.New(kafkaClient)
	topicController := tc.New(topicService)
	consumerGroupsController := cgc.New(consumerGroupsService)

	// 라우트 설정
	for _, route := range topicController.Routes() {
		r.Handle(route.Method, route.Path, gin.HandlerFunc(route.Handler))
	}
	for _, route := range consumerGroupsController.Routes() {
		r.Handle(route.Method, route.Path, gin.HandlerFunc(route.Handler))
	}

	// 정적 파일 서버 설정
	r.Static("/static", "./static")

	// 404 Not Found 핸들러
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "404 - Not Found"})
	})

	fmt.Println("Starting server at port 8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
