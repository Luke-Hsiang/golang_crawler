package news

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var group *gin.RouterGroup

// Init app
func Init(gormdb *gorm.DB, g *gin.RouterGroup) {
	db = gormdb
	group = g
	InitModel()
	InitRoute()
}
