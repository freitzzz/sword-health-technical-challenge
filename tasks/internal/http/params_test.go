package http

import (
	"testing"
)

func TestParsePaginationIndexIntegerAllowed(t *testing.T) {
	idx := "0"

	_, err := ParsePaginationIndex(idx)

	if err != nil {
		t.Fatalf("Failed to parse %s pagination index", idx)
	}

}

func TestParsePaginationIndexNonIntegerAllowed(t *testing.T) {
	idx := "a"

	_, err := ParsePaginationIndex(idx)

	if err == nil {
		t.Fatalf(`Pagination index "%s" is not an integer, but parse did not fail`, idx)
	}

}

func TestParsePaginationIndexReturnsCorrectValue(t *testing.T) {
	idx := "42"
	midx := 42

	pidx, _ := ParsePaginationIndex(idx)

	if pidx != midx {
		t.Fatalf(`Parse should have returned %d, but instead returned %d`, midx, pidx)
	}

}
