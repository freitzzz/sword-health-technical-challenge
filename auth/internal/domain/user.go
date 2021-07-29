package domain

import (
	"time"

	"gorm.io/gorm"
)

const (
	technician       = 0
	manager          = 1
	mockTechnicianId = "tech"
	mockManagerId    = "man"
	mockSecret       = "sword"
)

var (
	MockTechnician = NewTechnician(mockTechnicianId)
	MockManager    = NewManager(mockManagerId)
)

type User struct {
	Identifier string `gorm:"uniqueIndex"`
	Secret     string
	Role       int
	gorm.Model
}

type UserSession struct {
	UserID          string
	Token           string
	ExpireTimestamp int64
	gorm.Model
}

func NewTechnician(id string) User {

	return User{Identifier: id, Secret: mockSecret, Role: technician}
}

func NewManager(id string) User {

	return User{Identifier: id, Secret: mockSecret, Role: manager}
}

func ValidAuth(u User, uid string, s string) bool {
	return u.Identifier == uid && u.Secret == s
}

func NewSession(u User, b JWTBundle) (UserSession, error) {

	var us UserSession

	exp := time.Now().AddDate(0, 0, 3).Unix()

	jt, serr := SignUserSession(b, u, exp)

	if serr == nil {
		us = UserSession{UserID: u.Identifier, ExpireTimestamp: exp, Token: jt}
	}

	return us, serr

}
