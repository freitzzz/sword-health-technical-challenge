package domain

import (
	"strings"
	"testing"
)

func TestNewTaskWithSummaryThatExceeds2500Characters(t *testing.T) {
	uid := "x"
	summary := strings.Repeat("x", 2501)

	_, err := New(uid, summary)

	if err == nil {
		t.Fatalf("Summary designated by the string below exceeds 2500 characters, but no error was identified\nString: %s", summary)
	}

}

func TestNewTaskWithSummaryThatMatches2500Characters(t *testing.T) {
	uid := "x"
	summary := strings.Repeat("x", 2500)

	_, err := New(uid, summary)

	if err != nil {
		t.Fatalf("Summary designated by the string below matches 2500 characters, but an error was identified\nString: %s", summary)
	}

}

func TestNewTaskWithSummaryThatDoesNotExceed2500Characters(t *testing.T) {
	uid := "x"
	summary := strings.Repeat("x", 2499)

	_, err := New(uid, summary)

	if err != nil {
		t.Fatalf("Summary designated by the string below does not exceed 2500 characters, but an error was identified\nString: %s", summary)
	}

}

func TestNewTaskTrimsSpaces(t *testing.T) {
	uid := "x"
	summary := strings.Repeat(" ", 50)

	task, _ := New(uid, summary)

	if len(task.Summary) == len(summary) {
		t.Fatalf("Original summary string contains unneeded spaces, and should be trimmed, but output summary still contains those spaces")
	}

}
