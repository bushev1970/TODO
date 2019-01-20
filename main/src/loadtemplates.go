package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "main website",
	})
}

func Day(c *gin.Context) {
	c.HTML(http.StatusOK, "day.html", gin.H{
		"title": "main website",
	})
}

func Week(c *gin.Context) {
	c.HTML(http.StatusOK, "week.html", gin.H{
		"title": "main website",
	})
}

func Month(c *gin.Context) {
	c.HTML(http.StatusOK, "month.html", gin.H{
		"title": "main website",
	})
}

func Year(c *gin.Context) {
	c.HTML(http.StatusOK, "year.html", gin.H{
		"title": "main website",
	})
}

func Setting(c *gin.Context) {
	c.HTML(http.StatusOK, "setting.html", gin.H{
		"title": "main website",
	})
}
