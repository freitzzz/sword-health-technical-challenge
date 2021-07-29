package domain

import (
	"gorm.io/gorm"
)

type Notification struct {
	TechnicianUserID string
	Message          string
	gorm.Model
}

type NotificationRead struct {
	ManagerUserID  string
	NotificationID uint
	gorm.Model
}

func New(message string, uid string) Notification {

	return Notification{
		TechnicianUserID: uid,
		Message:          message,
	}

}

func MarkAsRead(notification *Notification, uid string) *NotificationRead {
	return &NotificationRead{
		ManagerUserID:  uid,
		NotificationID: notification.ID,
	}
}
