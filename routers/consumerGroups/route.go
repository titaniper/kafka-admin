package consumerGroups

import (
	"encoding/json"
	"github.com/titaniper/gopang/services/consumerGroups"
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

func (c *Controller) ConsumerGroupsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.Get(w, r)
	case http.MethodDelete:
		c.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *Controller) Routes() []struct {
	Path    string
	Handler http.HandlerFunc
} {
	return []struct {
		Path    string
		Handler http.HandlerFunc
	}{
		{Path: "/consumer-groups", Handler: c.ConsumerGroupsHandler},
	}
}
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.consumerGroupsService.Delete(req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	topics, err := c.consumerGroupsService.List(keyword)
	//json.NewEncoder(w).Encode(map[string][]string{
	//	"data": {"하이", "하이"},
	//})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(map[string][]string{
		"data": topics,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
