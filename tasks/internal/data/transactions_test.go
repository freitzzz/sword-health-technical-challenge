package data

import (
	"testing"
)

func TestPaginationIndexToOffsetMultipliesIndexByQueryLimit(t *testing.T) {
	pidx := 2
	expectedOffset := pidx * _queryResultsLimit

	offset := PaginationIndexToOffset(pidx)

	if offset != expectedOffset {
		t.Fatalf("Expected offset is %d, but got %d instead", expectedOffset, offset)
	}

}
