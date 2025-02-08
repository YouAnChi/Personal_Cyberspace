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

	// 先清空数据库中的文章
	if err := config.DB.Exec("DELETE FROM posts").Error; err != nil {
		return err
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

			// 创建新文章
			htmlContent := string(markdown.ToHTML(content, parser.NewWithExtensions(parser.CommonExtensions), html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})))
			post := models.Post{
				Title:       title,
				Content:     string(content),
				HTMLContent: htmlContent,
				Author:      "admin",
				IsPublished: true,
			}

			// 保存到数据库
			if err := config.DB.Create(&post).Error; err != nil {
				log.Printf("保存文章 %s 失败: %v", title, err)
				return nil
			}

			log.Printf("成功加载文章: %s", title)
		}
		return nil
	})
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
