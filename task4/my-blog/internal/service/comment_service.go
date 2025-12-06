package service

import (
	"blog-backend/internal/model"
	"gorm.io/gorm"
)

type CommentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{db: db}
}

func (s *CommentService) Create(userID uint, req *model.CreateCommentRequest) (*model.Comment, error) {
	var post model.Post
	if err := s.db.First(&post, req.PostID).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	comment := &model.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  req.PostID,
	}

	if err := s.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *CommentService) GetByID(id uint) (*model.Comment, error) {
	var comment model.Comment
	if err := s.db.Preload("User").First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (s *CommentService) ListByPostID(postID uint, limit, offset int) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64

	if err := s.db.Model(&model.Comment{}).Where("post_id = ?", postID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := s.db.Preload("User").Where("post_id = ?", postID).
		Limit(limit).Offset(offset).Order("created_at DESC").Find(&comments).Error; err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

func (s *CommentService) Update(userID, commentID uint, req *model.UpdateCommentRequest) (*model.Comment, error) {
	var comment model.Comment
	if err := s.db.Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error; err != nil {
		return nil, err
	}

	comment.Content = req.Content

	if err := s.db.Save(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (s *CommentService) Delete(userID, commentID uint) error {
	return s.db.Where("id = ? AND user_id = ?", commentID, userID).Delete(&model.Comment{}).Error
}

