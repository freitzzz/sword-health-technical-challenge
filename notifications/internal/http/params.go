package http

import "strconv"

const (
	notificationId = "id"
	userIdHeader   = "X-User-ID"
	roleHeader     = "X-User-Role"
)

func ParsePaginationIndex(idx string) (int, error) {
	return strconv.Atoi(idx)
}
