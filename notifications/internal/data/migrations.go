package data

import (
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/domain"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {

	return db.AutoMigrate(&domain.Notification{})

}
