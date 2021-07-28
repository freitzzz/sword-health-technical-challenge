package data

import (
	"gorm.io/gorm"

	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/domain"
)

var (
	readQueryMap = map[string]interface{}{"read": false}
)

func QueryUserNotifications(db *gorm.DB, uid string) []*domain.Notification {

	var notifications []*domain.Notification

	db.Where(&domain.Notification{UserID: uid}, "read").Find(&notifications)

	return notifications

}

func InsertNotification(db *gorm.DB, notification domain.Notification) (*domain.Notification, error) {

	result := db.Create(&notification)

	return &notification, result.Error

}

func QueryUserNotificationById(db *gorm.DB, uid string, tid int) (*domain.Notification, error) {

	var notification domain.Notification

	result := db.Where(&domain.Notification{UserID: uid}, "disabled").First(&notification, tid)

	return &notification, result.Error

}

func UpdateNotification(db *gorm.DB, notification domain.Notification) (*domain.Notification, error) {

	result := db.Save(&notification)

	return &notification, result.Error

}
