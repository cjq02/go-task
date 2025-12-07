package handler

import (
	"blog-backend/internal/model"
	"blog-backend/internal/response"
	"blog-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
	jwtService  *service.JWTService
}

func NewUserHandler(userService *service.UserService, jwtService *service.JWTService) *UserHandler {
	return &UserHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrCodeValidationFailed, err.Error())
		return
	}

	user, err := h.userService.Register(&req)
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, err.Error())
		return
	}

	response.SuccessWithMessage(c, user.ToResponse(), "注册成功")
}

func (h *UserHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, response.ErrCodeValidationFailed, err.Error())
		return
	}

	user, err := h.userService.Login(&req)
	if err != nil {
		response.Error(c, response.ErrCodeAuthFailed, err.Error())
		return
	}

	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, "生成令牌失败")
		return
	}

	response.SuccessWithMessage(c, gin.H{
		"user":  user.ToResponse(),
		"token": token,
	}, "登录成功")
}

func (h *UserHandler) Logout(c *gin.Context) {
	response.SuccessWithMessage(c, gin.H{}, "退出登录成功")
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, response.ErrCodeAuthFailed, "未找到用户信息")
		return
	}

	user, err := h.userService.GetByID(userID.(uint))
	if err != nil {
		response.Error(c, response.ErrCodeDatabaseError, "用户不存在")
		return
	}

	response.Success(c, user.ToResponse())
}
