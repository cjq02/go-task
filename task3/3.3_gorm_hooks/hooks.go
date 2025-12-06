package main

import (
	"fmt"
	"log"

	"go-task/task3/3.1_gorm_models/models"
)

func main() {
	db, err := models.GetDB()
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	err = models.AutoMigrate(db)
	if err != nil {
		log.Fatalf("迁移失败: %v", err)
	}

	user := models.User{
		Name:  "测试用户",
		Email: "test@example.com",
	}
	err = db.Create(&user).Error
	if err != nil {
		log.Fatalf("创建用户失败: %v", err)
	}

	post := models.Post{
		Title:   "测试文章",
		Content: "这是一篇测试文章",
		UserID:  user.ID,
	}
	err = db.Create(&post).Error
	if err != nil {
		log.Fatalf("创建文章失败: %v", err)
	}

	var updatedUser models.User
	db.First(&updatedUser, user.ID)
	fmt.Printf("用户文章数量: %d\n", updatedUser.PostCount)

	comment := models.Comment{
		Content: "这是一条评论",
		PostID:  post.ID,
	}
	err = db.Create(&comment).Error
	if err != nil {
		log.Fatalf("创建评论失败: %v", err)
	}

	err = db.Delete(&comment).Error
	if err != nil {
		log.Fatalf("删除评论失败: %v", err)
	}

	var updatedPost models.Post
	db.First(&updatedPost, post.ID)
	fmt.Printf("文章评论状态: %s\n", updatedPost.CommentStatus)
}
