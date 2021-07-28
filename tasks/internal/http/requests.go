package http

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type TaskPerform struct {
	Summary string `json:"summary"`
}

type TaskUpdate struct {
	Summary string `json:"summary"`
}

type UserContext struct {
	ID   string
	Role int
}

func IsTechnician(uc UserContext) bool {
	return uc.Role == 0
}
