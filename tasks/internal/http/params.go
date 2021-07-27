package http

import "strconv"

const (
	paginationIndex = "index"
	taskId          = "id"
)

func ParsePaginationIndex(idx string) (int, error) {
	return strconv.Atoi(idx)
}
