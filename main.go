package main

import (
	"log"
	"os"
	"os/signal"
	"pcy/config"
	"pcy/models"
	"pcy/routes"
	"syscall"
)

// 初始化测试数据
func initTestData() error {
	// 创建测试用户
	user := models.User{
		Username: "admin",
		Email:    "admin@example.com",
		Nickname: "管理员",
	}
	user.SetPassword("123456")

	// 使用新版GORM的FirstOrCreate
	result := config.DB.Where(models.User{Username: "admin"}).FirstOrCreate(&user)
	if result.Error != nil {
		return result.Error
	}

	// 创建测试文章
	posts := []models.Post{
		{
			Title:       "欢迎使用PCy个人网络空间",
			Content:     "这是一个使用Go语言和Bootstrap构建的个人网络空间系统。在这里，你可以分享你的想法、经验和故事。",
			Summary:     "系统介绍和使用说明",
			UserID:      user.ID,
			Category:    "公告",
			Tags:        "公告,使用说明",
			IsPublished: true,
		},
		{
			Title:       "Go语言入门教程",
			Content:     "Go是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。本文将介绍Go语言的基础知识和使用方法。",
			Summary:     "Go语言基础知识介绍",
			UserID:      user.ID,
			Category:    "技术",
			Tags:        "Go,编程,教程",
			IsPublished: true,
		},
	}

	for _, post := range posts {
		result := config.DB.Where(models.Post{Title: post.Title}).FirstOrCreate(&post)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func main() {
	// 初始化数据库连接
	if err := config.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 自动迁移数据库表
	log.Println("开始数据库迁移...")
	if err := config.DB.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	log.Println("数据库迁移完成")

	// 初始化测试数据
	log.Println("初始化测试数据...")
	if err := initTestData(); err != nil {
		log.Printf("初始化测试数据失败: %v", err)
	}
	log.Println("测试数据初始化完成")

	// 设置路由
	r := routes.SetupRouter()

	// 创建一个通道来接收系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 启动服务器
		if err := r.Run(":8080"); err != nil {
			log.Printf("服务器启动失败: %v", err)
			quit <- syscall.SIGTERM
		}
	}()

	// 等待退出信号
	<-quit
	log.Println("正在关闭服务器...")

	// 获取底层的sqlDB
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Printf("获取数据库连接失败: %v", err)
	} else {
		// 关闭数据库连接
		if err := sqlDB.Close(); err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
		} else {
			log.Println("数据库连接已成功关闭")
		}
	}

	log.Println("服务器已关闭")
}
