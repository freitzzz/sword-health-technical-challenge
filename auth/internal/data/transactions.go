package data

import (
	"gorm.io/gorm"

	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/domain"
)

func InsertUser(db *gorm.DB, u domain.User) (*domain.User, error) {

	result := db.Create(&u)

	return &u, result.Error

}

func InsertUserSession(db *gorm.DB, us domain.UserSession) (*domain.UserSession, error) {

	result := db.Create(&us)

	return &us, result.Error

}

func QueryUserByIdentifier(db *gorm.DB, uid string) (*domain.User, error) {

	var u domain.User

	result := db.First(&u, &domain.User{Identifier: uid})

	return &u, result.Error

}

func QueryUserSessionByToken(db *gorm.DB, tk string) (*domain.UserSession, error) {

	var us domain.UserSession

	result := db.First(&us, &domain.UserSession{Token: tk})

	return &us, result.Error

}
