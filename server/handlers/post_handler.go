package handlers

import (
	"go-assignment/models"
	"go-assignment/repositories"
	"go-assignment/requests"
	"go-assignment/responses"
	s "go-assignment/server"
	postservice "go-assignment/services/post"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandlers struct {
	server *s.Server
}

func NewPostHandlers(server *s.Server) *PostHandlers {
	return &PostHandlers{server: server}
}

func (p *PostHandlers) CreatePost(c *gin.Context) {
	createPostRequest := new(requests.CreatePostRequest)

	if err := c.ShouldBind(&createPostRequest); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := createPostRequest.Validate(); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
		return
	}
	user := c.GetString("user")
	userId, err := strconv.ParseUint(user, 10, 64) // base 10, up to 64 bits
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Invalid User Id")
		return
	}

	post := models.Post{
		Title:   createPostRequest.Title,
		Content: createPostRequest.Content,
		UserID:  uint(userId),
	}
	postService := postservice.NewPostService(p.server.DB)
	postService.Create(&post)

	responses.MessageResponse(c, http.StatusCreated, "Post successfully created")
	return
}

func (p *PostHandlers) DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	post := models.Post{}

	postRepository := repositories.NewPostRepository(p.server.DB)
	postRepository.GetPost(&post, id)
	if post.ID == 0 {
		responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	postService := postservice.NewPostService(p.server.DB)
	postService.Delete(&post)

	responses.MessageResponse(c, http.StatusNoContent, "Post deleted successfully")
	return
}

func (p *PostHandlers) GetPosts(c *gin.Context) {
	var posts []models.Post
	//user := c.GetUint("user")
	userId := c.GetString("user")

	postRepository := repositories.NewPostRepository(p.server.DB)
	postRepository.GetPostsForUser(&posts, userId)

	response := responses.NewPostResponse(posts)
	responses.Response(c, http.StatusOK, response)
	return
}

func (p *PostHandlers) UpdatePost(c *gin.Context) {
	updatePostRequest := new(requests.UpdatePostRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBind(&updatePostRequest); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := updatePostRequest.Validate(); err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
		return
	}

	post := models.Post{}

	postRepository := repositories.NewPostRepository(p.server.DB)
	postRepository.GetPost(&post, id)

	if post.ID == 0 {
		responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
		return
	}

	postService := postservice.NewPostService(p.server.DB)
	postService.Update(&post, updatePostRequest)

	responses.MessageResponse(c, http.StatusOK, "Post successfully updated")
	return
}
