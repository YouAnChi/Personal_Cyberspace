package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	// 数据库配置
	username := "root"   // 数据库用户名
	password := "123456" // 数据库密码
	host := "127.0.0.1"
	port := 3306
	dbName := "pcy"

	// 首先尝试创建数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 创建数据库
	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", dbName)
	err = db.Exec(createDB).Error
	if err != nil {
		return fmt.Errorf("创建数据库失败: %v", err)
	}

	// 连接到指定的数据库
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	return nil
}
