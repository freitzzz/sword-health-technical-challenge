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

type Task struct {
	ID                 string `json:"id"`
	UserID             string `json:"userId"`
	Summary            string `json:"summary"`
	CreatedTimestampMS int    `json:"createdTimestampMS"`
	UpdatedTimestampMS int    `json:"updatedTimestampMS"`
}