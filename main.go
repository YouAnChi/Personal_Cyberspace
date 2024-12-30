package main

import (
	"log"
	"os"
	"os/signal"
	"pcy/config"
	"pcy/models"
	"pcy/routes"
	"pcy/utils"
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
	user.SetPassword("12345678")

	// 使用新版GORM的FirstOrCreate
	result := config.DB.Where(models.User{Username: "admin"}).FirstOrCreate(&user)
	if result.Error != nil {
		return result.Error
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

	// 监控文章目录
	articleDir := "./articles/md"
	err := utils.LoadArticlesFromMarkdown(articleDir)
	if err != nil {
		log.Printf("初始文章加载失败: %v", err)
	}
	utils.WatchArticlesDirectory(articleDir)

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
