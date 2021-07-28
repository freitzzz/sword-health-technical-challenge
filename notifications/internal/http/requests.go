package http

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type UserContext struct {
	ID   string
	Role int
}

func IsTechnician(uc UserContext) bool {
	return uc.Role == 0
}
