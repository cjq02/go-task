package model

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	UserID    uint      `json:"userId" gorm:"not null"`
	PostID    uint      `json:"postId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Post      Post      `json:"-" gorm:"foreignKey:PostID"`
}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"postId" binding:"required"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

type CommentResponse struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	UserID    uint      `json:"userId"`
	PostID    uint      `json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
	User      *UserResponse `json:"user,omitempty"`
}

func (c *Comment) ToResponse() *CommentResponse {
	resp := &CommentResponse{
		ID:        c.ID,
		Content:   c.Content,
		UserID:    c.UserID,
		PostID:    c.PostID,
		CreatedAt: c.CreatedAt,
	}
	if c.User.ID != 0 {
		resp.User = c.User.ToResponse()
	}
	return resp
}

