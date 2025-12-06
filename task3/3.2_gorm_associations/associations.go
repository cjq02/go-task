package main

import (
	"fmt"
	"log"

	"go-task/task3/3.1_gorm_models/models"

	"gorm.io/gorm"
)

func queryUserPostsWithComments(db *gorm.DB, userID uint) (models.User, error) {
	var user models.User
	err := db.Preload("Posts.Comments").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return models.User{}, fmt.Errorf("查询用户文章及评论失败: %v", err)
	}
	return user, nil
}

func queryPostWithMostComments(db *gorm.DB) (models.Post, error) {
	var post models.Post
	var postIDs []uint

	err := db.Model(&models.Comment{}).
		Select("post_id").
		Group("post_id").
		Order("COUNT(*) DESC").
		Limit(1).
		Pluck("post_id", &postIDs).Error
	if err != nil {
		return models.Post{}, fmt.Errorf("查询评论数量最多的文章失败: %v", err)
	}

	if len(postIDs) == 0 {
		return models.Post{}, fmt.Errorf("没有找到有评论的文章")
	}

	err = db.Preload("Comments").Preload("User").First(&post, postIDs[0]).Error
	if err != nil {
		return models.Post{}, fmt.Errorf("查询文章详情失败: %v", err)
	}

	return post, nil
}

func main() {
	db, err := models.GetDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	userID := uint(1)
	user, err := queryUserPostsWithComments(db, userID)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	fmt.Printf("用户 %s 的所有文章及其评论：\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("  文章: %s (评论数: %d)\n", post.Title, len(post.Comments))
	}

	post, err := queryPostWithMostComments(db)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}

	fmt.Printf("评论数量最多的文章: %s (评论数: %d)\n", post.Title, len(post.Comments))
}
