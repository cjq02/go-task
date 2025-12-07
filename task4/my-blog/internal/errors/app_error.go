package errors

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code       int    // 业务错误码
	HTTPStatus int    // HTTP 状态码
	Message    string // 错误消息
	Err        error  // 原始错误
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// 错误码定义
const (
	ErrCodeValidationFailed = 1001 // 参数验证失败
	ErrCodeAuthFailed       = 1002 // 认证失败
	ErrCodeForbidden        = 1003 // 无权限
	ErrCodeNotFound         = 1004 // 资源不存在
	ErrCodeDatabaseError    = 2001 // 数据库错误
	ErrCodeInternalError    = 2002 // 内部服务器错误
)

// 创建错误
func NewValidationError(message string) *AppError {
	return &AppError{
		Code:       ErrCodeValidationFailed,
		HTTPStatus: http.StatusBadRequest,
		Message:    message,
	}
}

func NewAuthError(message string) *AppError {
	return &AppError{
		Code:       ErrCodeAuthFailed,
		HTTPStatus: http.StatusUnauthorized,
		Message:    message,
	}
}

func NewForbiddenError(message string) *AppError {
	return &AppError{
		Code:       ErrCodeForbidden,
		HTTPStatus: http.StatusForbidden,
		Message:    message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:       ErrCodeNotFound,
		HTTPStatus: http.StatusNotFound,
		Message:    message,
	}
}

func NewDatabaseError(err error) *AppError {
	return &AppError{
		Code:       ErrCodeDatabaseError,
		HTTPStatus: http.StatusInternalServerError,
		Message:    "数据库操作失败",
		Err:        err,
	}
}

func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Code:       ErrCodeInternalError,
		HTTPStatus: http.StatusInternalServerError,
		Message:    message,
		Err:        err,
	}
}

// 判断是否为 AppError
func IsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

// 从标准错误转换为 AppError
func WrapError(err error) *AppError {
	if err == nil {
		return nil
	}

	if appErr, ok := IsAppError(err); ok {
		return appErr
	}

	// 默认作为数据库错误处理
	return NewDatabaseError(err)
}
