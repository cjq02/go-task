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

type PostHandler struct {
	postService *service.PostService
	logger      *logger.Logger
}

func NewPostHandler(postService *service.PostService, appLogger *logger.Logger) *PostHandler {
	return &PostHandler{
		postService: postService,
		logger:      appLogger,
	}
}

func (h *PostHandler) Create(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)

	var req model.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewValidationError("请求参数无效: " + err.Error())
		response.HandleError(c, h.logger, appErr)
		return
	}

	post, err := h.postService.Create(userID, &req)
	if err != nil {
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.Info("用户 %d 创建了文章，ID: %d", userID, post.ID)
	response.Success(c, post.ToResponse())
}

func (h *PostHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appErr := errors.NewValidationError("无效的文章ID")
		response.HandleError(c, h.logger, appErr)
		return
	}

	post, err := h.postService.GetByID(uint(id))
	if err != nil {
		response.HandleGormError(c, h.logger, err, "文章")
		return
	}

	response.Success(c, post.ToResponse())
}

func (h *PostHandler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	posts, total, err := h.postService.List(limit, offset)
	if err != nil {
		response.HandleError(c, h.logger, err)
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
		appErr := errors.NewValidationError("无效的文章ID")
		response.HandleError(c, h.logger, appErr)
		return
	}

	var req model.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appErr := errors.NewValidationError("请求参数无效: " + err.Error())
		response.HandleError(c, h.logger, appErr)
		return
	}

	post, err := h.postService.Update(userID, uint(id), &req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			appErr := errors.NewForbiddenError("无权限修改该文章或文章不存在")
			response.HandleError(c, h.logger, appErr)
			return
		}
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.Info("用户 %d 更新了文章，ID: %d", userID, post.ID)
	response.Success(c, post.ToResponse())
}

func (h *PostHandler) Delete(c *gin.Context) {
	userIDInterface, _ := c.Get("userID")
	userID := userIDInterface.(uint)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appErr := errors.NewValidationError("无效的文章ID")
		response.HandleError(c, h.logger, appErr)
		return
	}

	if err := h.postService.Delete(userID, uint(id)); err != nil {
		if err == gorm.ErrRecordNotFound {
			appErr := errors.NewForbiddenError("无权限删除该文章或文章不存在")
			response.HandleError(c, h.logger, appErr)
			return
		}
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.Info("用户 %d 删除了文章，ID: %d", userID, id)
	response.Success(c, gin.H{"message": "删除成功"})
}
