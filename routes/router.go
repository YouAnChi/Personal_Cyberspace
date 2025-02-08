package routes

import (
	"net/http"
	"pcy/controllers"
	"pcy/metrics"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 添加指标收集中间件
	r.Use(metrics.MetricsMiddleware())

	// 添加Prometheus指标数据暴露端点
	r.GET("/metrics", metrics.PrometheusHandler())

	// 静态文件服务
	r.Static("/static", "./web/static")
	r.StaticFile("/favicon.ico", "./web/static/favicon.ico")
	r.LoadHTMLGlob("web/templates/*")

	// 设置信任的代理
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 前端页面路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cover.html", nil)
	})

	r.GET("/cover", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cover.html", nil)
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/blog", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", nil)
	})

	r.GET("/blog/post/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	// API路由组
	api := r.Group("/api")
	{
		// 用户相关
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// 文章相关
		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPost)
		api.POST("/posts", controllers.CreatePost)
		api.PUT("/posts/:id", controllers.UpdatePost)
		api.DELETE("/posts/:id", controllers.DeletePost)
	}

	return r
}
