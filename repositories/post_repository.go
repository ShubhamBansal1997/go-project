package repositories

import (
	"go-assignment/models"

	"gorm.io/gorm"
)

type PostRepositoryQ interface {
	GetPosts(posts *[]models.Post)
	GetPost(post *models.Post, id int)
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (postRepository *PostRepository) GetPostsForUser(posts *[]models.Post, userId string) {
	postRepository.DB.Where("user_id = ? ", userId).Find(posts)
}

func (postRepository *PostRepository) GetPost(post *models.Post, id int) {
	postRepository.DB.Where("id = ? ", id).First(post)
}
