package handler

import (
	"sync"
	"web-forum/internal/service"
	"web-forum/models"
)

var (
	userUrlBefore = make(map[string]string)
	mu            sync.Mutex
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func getTags(tags []string, form []models.GetAllPosts) []models.GetAllPosts {
	foundData := []models.GetAllPosts{}

	searchMap := make(map[string]bool)

	for _, tag := range tags {
		searchMap[tag] = true
	}
	// ds
	for _, post := range form {
		for _, postTag := range post.Tags {
			if searchMap[postTag] {
				foundData = append(foundData, post)
				break
			}
		}
	}
	return foundData
}
