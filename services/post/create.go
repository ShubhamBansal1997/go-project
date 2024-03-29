package post

import "go-assignment/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
