package data

import (
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/domain"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {

	return db.AutoMigrate(&domain.Task{})

}
