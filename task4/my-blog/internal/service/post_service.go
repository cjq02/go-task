package service

import (
	"blog-backend/internal/model"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) Create(userID uint, req *model.CreatePostRequest) (*model.Post, error) {
	post := &model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := s.db.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) GetByID(id uint) (*model.Post, error) {
	var post model.Post
	if err := s.db.Preload("User").Preload("Comments.User").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *PostService) List(limit, offset int) ([]model.Post, int64, error) {
	var posts []model.Post
	var total int64

	if err := s.db.Model(&model.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := s.db.Preload("User").Limit(limit).Offset(offset).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (s *PostService) Update(userID, postID uint, req *model.UpdatePostRequest) (*model.Post, error) {
	var post model.Post
	if err := s.db.Where("id = ? AND user_id = ?", postID, userID).First(&post).Error; err != nil {
		return nil, err
	}

	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	if err := s.db.Save(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostService) Delete(userID, postID uint) error {
	return s.db.Where("id = ? AND user_id = ?", postID, userID).Delete(&model.Post{}).Error
}

