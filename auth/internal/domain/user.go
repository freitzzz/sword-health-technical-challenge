package domain

import (
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
	Identifier string
	Secret     string
	Role       int
	gorm.Model
}

type UserSession struct {
	UserID          string
	Token           string
	ExpireTimestamp uint64
	gorm.Model
}

func NewTechnician(id string) User {

	return User{Identifier: id, Secret: mockSecret, Role: technician}
}

func NewManager(id string) User {

	return User{Identifier: id, Secret: mockSecret, Role: manager}
}

func NewSession(u User) UserSession {
	return UserSession{UserID: u, Token: }
}