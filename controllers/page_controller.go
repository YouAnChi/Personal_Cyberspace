package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RenderCover renders the cover page
func RenderCover(c *gin.Context) {
	c.HTML(http.StatusOK, "cover.html", gin.H{})
}

// RenderHome renders the home page
func RenderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}
