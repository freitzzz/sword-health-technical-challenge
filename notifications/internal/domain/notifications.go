package domain

import (
	"gorm.io/gorm"
)

type Notification struct {
	UserID  string
	Message string
	gorm.Model
}

type NotificationRead struct {
	UserID         string
	NotificationID uint
	gorm.Model
}

func New(message string, uid string) Notification {

	return Notification{
		UserID:  uid,
		Message: message,
	}

}

func MarkAsRead(notification *Notification, uid string) *NotificationRead {
	return &NotificationRead{
		UserID:         uid,
		NotificationID: notification.ID,
	}
}
