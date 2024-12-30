package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"pcy/config"
	"pcy/models"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// LoadArticlesFromMarkdown 从 Markdown 文件加载文章
func LoadArticlesFromMarkdown(mdDir string) error {
	// 确保目录存在
	if _, err := os.Stat(mdDir); os.IsNotExist(err) {
		os.MkdirAll(mdDir, 0755)
	}

	// 遍历目录中的 Markdown 文件
	return filepath.Walk(mdDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理 .md 文件
		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			// 读取文件内容
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("读取文件 %s 失败: %v", path, err)
				return nil
			}

			// 解析文件名作为文章信息
			filename := filepath.Base(path)
			title := strings.TrimSuffix(filename, ".md")

			// 检查文章是否已存在
			var existingPost models.Post
			result := config.DB.Where("title = ?", title).First(&existingPost)
			if result.Error == nil {
				log.Printf("文章 %s 已存在，跳过", title)
				return nil
			}

			// 更全面的 Markdown 扩展
			extensions := parser.CommonExtensions |
				parser.AutoHeadingIDs |
				parser.NoEmptyLineBeforeBlock |
				parser.HardLineBreak |      // 启用硬换行
				parser.Strikethrough        // 删除线

			mdParser := parser.NewWithExtensions(extensions)

			htmlFlags := html.CommonFlags |
				html.HrefTargetBlank |
				html.CompletePage | // 生成完整的 HTML 页面
				html.UseXHTML // 使用 XHTML 标准

			opts := html.RendererOptions{
				Flags: htmlFlags,
				CSS:   "", // 可以添加自定义 CSS
				Title: title,
			}
			renderer := html.NewRenderer(opts)

			htmlContent := string(markdown.ToHTML(content, mdParser, renderer))

			// 创建文章对象
			post := models.Post{
				Title:       title,
				Content:     string(content),
				Summary:     generateSummary(htmlContent),
				Category:    "文章",
				Tags:        "markdown",
				IsPublished: true,
				UserID:      1, // 默认管理员
			}

			// 保存到数据库
			result = config.DB.Create(&post)
			if result.Error != nil {
				log.Printf("保存文章 %s 失败: %v", title, result.Error)
			} else {
				log.Printf("成功导入文章：%s", title)
			}
		}
		return nil
	})
}

// generateSummary 从 HTML 内容生成摘要
func generateSummary(htmlContent string) string {
	// 简单实现：取前200个字符
	runes := []rune(htmlContent)
	if len(runes) > 200 {
		return string(runes[:200]) + "..."
	}
	return string(runes)
}

// WatchArticlesDirectory 监控文章目录
func WatchArticlesDirectory(mdDir string) {
	go func() {
		for {
			err := LoadArticlesFromMarkdown(mdDir)
			if err != nil {
				log.Printf("加载文章失败: %v", err)
			}
			// 每小时检查一次
			time.Sleep(1 * time.Hour)
		}
	}()
}
