package http

type InvalidParam struct {
	Message string
	Type    string
	Name    string
}

var (
	paginationIndexNotInteger = InvalidParam{
		Message: "Not an integer",
		Type:    "Query",
		Name:    paginationIndex,
	}
)
