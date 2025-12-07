package handler

import (
	"blog-backend/internal/errors"
	"blog-backend/internal/logger"
	"blog-backend/internal/model"
	"blog-backend/internal/response"
	"blog-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentHandler struct {
	commentService *service.CommentService
	logger         *logger.Logger
}

func NewCommentHandler(commentService *service.CommentService, appLogger *logger.Logger) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
		logger:         appLogger,
	}
}

func (h *CommentHandler) Create(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)

	var req model.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewValidationError("请求参数无效: " + err.Error())
		response.HandleError(c, h.logger, appErr)
		return
	}

	comment, err := h.commentService.Create(userID, &req)
	if err != nil {
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.Info("用户 %d 创建了评论，ID: %d, 文章ID: %d", userID, comment.ID, comment.PostID)
	response.Success(c, comment.ToResponse())
}

func (h *CommentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appErr := errors.NewValidationError("无效的评论ID")
		response.HandleError(c, h.logger, appErr)
		return
	}

	comment, err := h.commentService.GetByID(uint(id))
	if err != nil {
		response.HandleGormError(c, h.logger, err, "评论")
		return
	}

	response.Success(c, comment.ToResponse())
}

func (h *CommentHandler) ListByPostID(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("postId"), 10, 32)
	if err != nil {
		appErr := errors.NewValidationError("无效的文章ID")
		response.HandleError(c, h.logger, appErr)
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	comments, total, err := h.commentService.ListByPostID(uint(postID), limit, offset)
	if err != nil {
		response.HandleError(c, h.logger, err)
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
		appErr := errors.NewValidationError("无效的评论ID")
		response.HandleError(c, h.logger, appErr)
		return
	}

	var req model.UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewValidationError("请求参数无效: " + err.Error())
		response.HandleError(c, h.logger, appErr)
		return
	}

	comment, err := h.commentService.Update(userID, uint(id), &req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appErr := errors.NewForbiddenError("无权限修改该评论或评论不存在")
			response.HandleError(c, h.logger, appErr)
			return
		}
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.Info("用户 %d 更新了评论，ID: %d", userID, comment.ID)
	response.Success(c, comment.ToResponse())
}

func (h *CommentHandler) Delete(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appErr := errors.NewValidationError("无效的评论ID")
		response.HandleError(c, h.logger, appErr)
		return
	}

	if err := h.commentService.Delete(userID, uint(id)); err != nil {
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.Info("用户 %d 删除了评论，ID: %d", userID, id)
	response.Success(c, gin.H{"message": "删除成功"})
}
