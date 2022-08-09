package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRoute init route
func InitRoute() {
	news := group.Group("/news")
	news.GET("/news", getAllNews)
	news.DELETE("/news", deleteAllNews)
	news.POST("/reporters", createReporter)
	news.GET("/reporters", getAllReporters)
	news.DELETE("/reporters", deleteAllReporters)
}

func getAllNews(c *gin.Context) {
	var result []News
	db.Preload("Reporter").Find(&result)
	c.JSON(http.StatusOK, result)
}

func deleteAllNews(c *gin.Context) {
	db.Unscoped().Where("1 = 1").Delete(&News{})
	c.JSON(http.StatusOK, "")
}

func createReporter(c *gin.Context) {
	reporter := Reporter{Name: "Luke"}
	db.Create(&reporter)
	c.JSON(http.StatusOK, reporter)
}

func getAllReporters(c *gin.Context) {
	var result []Reporter
	db.Find(&result)
	c.JSON(http.StatusOK, result)
}

func deleteAllReporters(c *gin.Context) {
	result := db.Unscoped().Where("1 = 1").Delete(&Reporter{})
	c.JSON(http.StatusOK, result.Error)
}
