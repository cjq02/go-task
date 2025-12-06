package main

import (
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
		log.Fatalf("创建表失败: %v", err)
	}

	log.Println("模型表创建成功！")
}

