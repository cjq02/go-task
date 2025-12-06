package handler

import (
	"blog-backend/internal/model"
	"blog-backend/internal/response"
	"blog-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) Create(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)

	var req model.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrCodeValidationFailed, err.Error())
		return
	}

	comment, err := h.commentService.Create(userID, &req)
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	response.Success(c, comment.ToResponse())
}

func (h *CommentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrCodeValidationFailed, "无效的评论ID")
		return
	}

	comment, err := h.commentService.GetByID(uint(id))
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, "评论不存在")
		return
	}

	response.Success(c, comment.ToResponse())
}

func (h *CommentHandler) ListByPostID(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("postId"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrCodeValidationFailed, "无效的文章ID")
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	comments, total, err := h.commentService.ListByPostID(uint(postID), limit, offset)
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	var commentResponses []*model.CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, comment.ToResponse())
	}

	response.Success(c, gin.H{
		"list":  commentResponses,
		"total": total,
	})
}

func (h *CommentHandler) Update(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrCodeValidationFailed, "无效的评论ID")
		return
	}

	var req model.UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrCodeValidationFailed, err.Error())
		return
	}

	comment, err := h.commentService.Update(userID, uint(id), &req)
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	response.Success(c, comment.ToResponse())
}

func (h *CommentHandler) Delete(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, response.ErrCodeValidationFailed, "无效的评论ID")
		return
	}

	if err := h.commentService.Delete(userID, uint(id)); err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

