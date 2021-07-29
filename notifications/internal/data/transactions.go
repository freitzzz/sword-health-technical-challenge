package data

import (
	"errors"

	"gorm.io/gorm"

	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/domain"
)

func QueryUserNotifications(db *gorm.DB, uid string) []*domain.Notification {

	var notifications []*domain.Notification

	var notificationsReadIds []uint

	nrq := &domain.NotificationRead{UserID: uid}

	db.Model(&nrq).Where(&nrq).Select("notification_id").Find(&notificationsReadIds)

	db.Not(notificationsReadIds).Find(&notifications)

	return notifications

}

func InsertNotification(db *gorm.DB, notification domain.Notification) (*domain.Notification, error) {

	result := db.Create(&notification)

	return &notification, result.Error

}

func InsertNotificationRead(db *gorm.DB, notificationRead domain.NotificationRead) (*domain.NotificationRead, error) {

	result := db.Create(&notificationRead)

	return &notificationRead, result.Error

}

func QueryUserNotificationById(db *gorm.DB, uid string, nid uint) (*domain.Notification, error) {

	var notification domain.Notification

	var notificationsReadIds []uint

	nrq := &domain.NotificationRead{UserID: uid, NotificationID: nid}

	db.Model(&nrq).Where(&nrq).Select("id").Find(&notificationsReadIds)

	if len(notificationsReadIds) != 0 {
		return nil, errors.New("User has already marked notification as read")
	}

	result := db.First(&notification, nid)

	return &notification, result.Error

}
