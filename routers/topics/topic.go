package topics

import (
	"encoding/json"
	"github.com/titaniper/gopang/services/topics"
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

func (c *Controller) TopicsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.Create(w, r)
	case http.MethodGet:
		c.Get(w, r)
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
		{Path: "/topics", Handler: c.TopicsHandler},
	}
}
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.topicService.CreateTopic(req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	topics, err := c.topicService.GetTopics(keyword)
	//json.NewEncoder(w).Encode(map[string][]string{
	//	"data": {"하이", "하이"},
	//})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(topics); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
