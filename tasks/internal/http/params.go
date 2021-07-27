package http

import "strconv"

const (
	paginationIndex = "index"
)

func ParsePaginationIndex(idx string) (int, error) {
	return strconv.Atoi(idx)
}
