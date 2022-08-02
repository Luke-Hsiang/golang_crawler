package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoute init route
func InitRoute() {
	news := group.Group("/news")
	news.GET("/news", getAllNews)
}

func getAllNews(c *gin.Context) {
	var result []News
	db.Find(&result)
	c.JSON(http.StatusOK, result)
}
