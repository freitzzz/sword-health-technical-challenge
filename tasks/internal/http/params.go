package http

import "strconv"

const (
	paginationIndex = "index"
	summaryField    = "summary"
	taskId          = "id"
	userIdHeader    = "X-User-ID"
	roleHeader      = "X-User-Role"
)

func ParsePaginationIndex(idx string) (int, error) {
	return strconv.Atoi(idx)
}
