package domain

import (
	"testing"
	"time"
)

func TestNewTechnicianHasTechnicianRole(t *testing.T) {
	u := NewTechnician(mockTechnicianId)

	if u.Role != technician {
		t.Fatalf("Technician created with role that does not identify a technician")
	}

}

func TestNewManagerHasManagerRole(t *testing.T) {
	u := NewManager(mockManagerId)

	if u.Role != manager {
		t.Fatalf("Manager created with role that does not identify a manager")
	}

}

func TestValidAuthReturnsTrueIfBothUserIDAndSecretMatch(t *testing.T) {
	u := NewManager(mockManagerId)

	va := ValidAuth(u, u.Identifier, u.Secret)

	if !va {
		t.Fatalf("User identifier and secret were passed to valid auth, but function returned false")
	}

}

func TestValidAuthReturnsFalseIfUserIDDoesNotMatch(t *testing.T) {
	u := NewManager(mockManagerId)

	va := ValidAuth(u, "abc", u.Secret)

	if va {
		t.Fatalf("Diferent user id was passed to valid auth, but function returned true")
	}

}

func TestValidAuthReturnsFalseIfSecretDoesNotMatch(t *testing.T) {
	u := NewManager(mockManagerId)

	va := ValidAuth(u, u.Identifier, "abc")

	if va {
		t.Fatalf("Diferent secret was passed to valid auth, but function returned true")
	}

}

func TestNewSessionReturnsUserSessionModelWithExpirationTimestampNoLongerThan3Days(t *testing.T) {
	u := NewManager(mockManagerId)

	db := time.Now()

	b := JWTBundle{Alg: "HS256", Secret: "asmalltinykeythatisultrasecureee"}

	us, _ := NewSession(u, b)

	da := time.Unix(us.ExpireTimestamp, 0)

	if db.After(da) {
		t.Fatalf("Session expiration date cannot be before current date")
	}

}
