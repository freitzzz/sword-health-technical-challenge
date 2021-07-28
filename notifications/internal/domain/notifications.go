package domain

import (
	"crypto/cipher"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Notification struct {
	UserID  string
	Message string
	Read    bool
	gorm.Model
}

func New(uid string, message string) Notification {

	return Notification{
		UserID:   uid,
		Message:  message
		Read: false,
	}

}

func MarkAsRed(notification *Notification) {
	notification.Read = true
}

