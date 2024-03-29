package post

import "go-assignment/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
