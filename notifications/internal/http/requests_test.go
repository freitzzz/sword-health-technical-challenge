package http

import "testing"

func TestIsTechnicianReturnsTrueIfRoleValueIs0(t *testing.T) {
	id := "x"
	role := 0

	uc := UserContext{ID: id, Role: role}

	check := IsTechnician(uc)

	if !check {
		t.Fatalf("User role value is %d, which is equal to 0, but IsTechnician call returned false", role)
	}

}

func TestIsTechnicianReturnsFalseIfRoleValueIs1(t *testing.T) {
	id := "x"
	role := 1

	uc := UserContext{ID: id, Role: role}

	check := IsTechnician(uc)

	if check {
		t.Fatalf("User role value is %d, which is equal to 1, but IsTechnician call returned true", role)
	}

}
