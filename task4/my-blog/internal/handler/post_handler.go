package handler

import (
	"blog-backend/internal/model"
	"blog-backend/internal/response"
	"blog-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) Create(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)

	var req model.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrCodeValidationFailed, err.Error())
		return
	}

	post, err := h.postService.Create(userID, &req)
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	response.Success(c, post.ToResponse())
}

func (h *PostHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrCodeValidationFailed, "无效的文章ID")
		return
	}

	post, err := h.postService.GetByID(uint(id))
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, "文章不存在")
		return
	}

	response.Success(c, post.ToResponse())
}

func (h *PostHandler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	posts, total, err := h.postService.List(limit, offset)
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	var postResponses []*model.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, post.ToResponse())
	}

	response.Success(c, gin.H{
		"list":  postResponses,
		"total": total,
	})
}

func (h *PostHandler) Update(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrCodeValidationFailed, "无效的文章ID")
		return
	}

	var req model.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrCodeValidationFailed, err.Error())
		return
	}

	post, err := h.postService.Update(userID, uint(id), &req)
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	response.Success(c, post.ToResponse())
}

func (h *PostHandler) Delete(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrCodeValidationFailed, "无效的文章ID")
		return
	}

	if err := h.postService.Delete(userID, uint(id)); err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

