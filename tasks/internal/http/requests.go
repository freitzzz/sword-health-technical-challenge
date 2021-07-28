package http

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type TaskPerform struct {
	Summary string `json:"summary"`
}

type TaskUpdate struct {
	Summary string `json:"summary"`
}

type TaskPage []struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}

type TaskView struct {
	ID                 string `json:"id"`
	UserID             string `json:"userId"`
	Summary            string `json:"summary"`
	CreatedTimestampMS int64  `json:"createdTimestampMS"`
	UpdatedTimestampMS int64  `json:"updatedTimestampMS"`
}

type UserContext struct {
	ID   string
	Role int
}

func IsTechnician(uc UserContext) bool {
	return uc.Role == 0
}
