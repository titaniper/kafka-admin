package topics

import (
	"github.com/gin-gonic/gin"
	"github.com/titaniper/kafka-admin/internal/services/topics"
	"net/http"
)

type Controller struct {
	topicService *topics.Service
}

func New(service *topics.Service) *Controller {
	return &Controller{
		topicService: service,
	}
}

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func (c *Controller) Routes() []Route {
	return []Route{
		{Method: http.MethodPost, Path: "/topics", Handler: c.Create},
		{Method: http.MethodGet, Path: "/topics", Handler: c.Get},
	}
}

// CreateTopicRequest represents the request structure for creating a topic
type CreateTopicRequest struct {
	Name string `json:"name" binding:"required"`
}

// TopicsResponse represents the response structure for topics
type TopicsResponse struct {
	Data []string `json:"data"`
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Create a new topic
// @Description Create a new Kafka topic
// @Tags topics
// @Accept json
// @Produce json
// @Param topic body CreateTopicRequest true "Topic Name"
// @Success 201 "Created"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /topics [post]
func (c *Controller) Create(ctx *gin.Context) {
	var req CreateTopicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := c.topicService.CreateTopic(req.Name); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

// @Summary Get topics
// @Description Get a list of Kafka topics
// @Tags topics
// @Accept json
// @Produce json
// @Param keyword query string false "Keyword to filter topics"
// @Success 200 {object} TopicsResponse
// @Failure 500 {object} ErrorResponse
// @Router /topics [get]
func (c *Controller) Get(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	topics, err := c.topicService.GetTopics(keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, TopicsResponse{Data: topics})
}
