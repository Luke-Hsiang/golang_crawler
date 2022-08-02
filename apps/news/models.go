package news

import (
	"gorm.io/gorm"
)

// Reporter is a person who write a news
type Reporter struct {
	gorm.Model
	Name string
}

// News is an article about recent events wrote by reporters
type News struct {
	gorm.Model
	Title   string
	Content string
}

// InitModel init models in current app
func InitModel() {
	db.AutoMigrate(&Reporter{})
	db.AutoMigrate(&News{})
}
