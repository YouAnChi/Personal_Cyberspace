package controllers

import (
	"log"
	"net/http"
	"pcy/config"
	"pcy/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePostRequest 创建文章请求结构
type CreatePostRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Summary     string `json:"summary"`
	Cover       string `json:"cover"`
	Category    string `json:"category"`
	Tags        string `json:"tags"`
	IsPublished bool   `json:"is_published"`
}

// GetPosts 获取文章列表
func GetPosts(c *gin.Context) {
	var posts []models.Post
	query := config.DB

	// 分页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	var total int64
	query.Model(&models.Post{}).Count(&total)

	if err := query.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts":     posts,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// GetPost 获取单篇文章
func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "文章ID不能为空",
			"code":  "INVALID_ID",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的文章ID格式",
			"code":  "INVALID_ID_FORMAT",
		})
		return
	}

	var post models.Post
	result := config.DB.First(&post, id)
	if result.Error != nil {
		// 尝试使用标题查找
		result = config.DB.Where("title = ?", idStr).First(&post)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "文章不存在或已被删除",
				"code":  "POST_NOT_FOUND",
			})
			return
		}
	}

	// 增加浏览次数（使用事务确保原子性）
	tx := config.DB.Begin()
	if err := tx.Model(&post).UpdateColumn("view_count", post.ViewCount+1).Error; err != nil {
		tx.Rollback()
		// 更新浏览次数失败不影响文章显示
		log.Printf("更新文章(%d)浏览次数失败: %v", post.ID, err)
	} else {
		tx.Commit()
		post.ViewCount++ // 更新返回对象的浏览次数
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
		"code": "SUCCESS",
	})
}

// CreatePost 创建文章
func CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效"})
		return
	}

	post := models.Post{
		Title:       req.Title,
		Content:     req.Content,
		Summary:     req.Summary,
		Cover:       req.Cover,
		Category:    req.Category,
		Tags:        req.Tags,
		Author:      "admin",
		IsPublished: req.IsPublished,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"post":    post,
	})
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数无效"})
		return
	}

	updates := models.Post{
		Title:       req.Title,
		Content:     req.Content,
		Summary:     req.Summary,
		Cover:       req.Cover,
		Category:    req.Category,
		Tags:        req.Tags,
		IsPublished: req.IsPublished,
	}

	if err := config.DB.Model(&post).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新文章失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"post":    post,
	})
}

// DeletePost 删除文章
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文章失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
