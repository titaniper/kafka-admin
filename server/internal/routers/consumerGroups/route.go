package consumerGroups

import (
	"github.com/gin-gonic/gin"
	"github.com/titaniper/kafka-admin/internal/services/consumerGroups"
	"net/http"
)

type Controller struct {
	consumerGroupsService *consumerGroups.Service
}

func New(service *consumerGroups.Service) *Controller {
	return &Controller{
		consumerGroupsService: service,
	}
}

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func (c *Controller) Routes() []Route {
	return []Route{
		{Method: http.MethodGet, Path: "/consumer-groups", Handler: c.Get},
		{Method: http.MethodDelete, Path: "/consumer-groups", Handler: c.Delete},
	}
}

// ConsumerGroupsResponse represents the response structure for consumer groups
type ConsumerGroupsResponse struct {
	Data []string `json:"data"`
}

// ErrorResponse represents the error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}

// DeleteRequest represents the request structure for deleting a consumer group
type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}

// @Summary Delete a consumer group
// @Description Delete a Kafka consumer group
// @Tags consumer-groups
// @Accept json
// @Produce json
// @Param group body DeleteRequest true "Consumer Group Name"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /consumer-groups [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	var req DeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := c.consumerGroupsService.Delete(req.Name); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// @Summary Get consumer groups
// @Description Get a list of Kafka consumer groups
// @Tags consumer-groups
// @Accept json
// @Produce json
// @Param keyword query string false "Keyword to filter consumer groups"
// @Success 200 {object} ConsumerGroupsResponse
// @Failure 500 {object} ErrorResponse
// @Router /consumer-groups [get]
func (c *Controller) Get(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	groups, err := c.consumerGroupsService.List(keyword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ConsumerGroupsResponse{Data: groups})
}
