package migration

import (
	"gorm.io/gorm"
	"task/app/model"
)

func Migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&model.Task{})
}
