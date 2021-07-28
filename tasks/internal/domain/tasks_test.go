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

func TestDisableMarksDisablePropertyAsTrue(t *testing.T) {
	uid := "x"
	summary := "x"

	task, _ := New(uid, summary)

	disabledPropBefore := task.Disabled

	Disable(&task)

	disabledPropAfter := task.Disabled

	if disabledPropBefore == disabledPropAfter {
		t.Fatalf("Disable properties before and after should be different")
	}

	if !disabledPropAfter {
		t.Fatalf("Disable procedure should mark Disabled property as false, but condition is not met")
	}

}

func TestUpdateSummaryWithStringThatExceeds2500Characters(t *testing.T) {
	uid := "x"
	summary := "x"

	task, _ := New(uid, summary)

	updatedSummary := strings.Repeat("x", 2501)

	err := UpdateSummary(&task, updatedSummary)

	if err == nil {
		t.Fatalf("Summary designated by the string below exceeds 2500 characters, but no error was identified on summary update\nString: %s", updatedSummary)
	}

}

func TestUpdateSummaryWithStringThatMatches2500Characters(t *testing.T) {
	uid := "x"
	summary := "x"

	task, _ := New(uid, summary)

	updatedSummary := strings.Repeat("x", 2500)

	err := UpdateSummary(&task, updatedSummary)

	if err != nil {
		t.Fatalf("Summary designated by the string below matches 2500 characters, but an error was identified on summary update\nString: %s", updatedSummary)
	}

}

func TestUpdateSummaryWithStringThatDoesNotExceed2500Characters(t *testing.T) {
	uid := "x"
	summary := "x"

	task, _ := New(uid, summary)

	updatedSummary := strings.Repeat("x", 2499)

	err := UpdateSummary(&task, updatedSummary)

	if err != nil {
		t.Fatalf("Summary designated by the string below does not exceed 2500 characters, but an error was identified on summary update\nString: %s", updatedSummary)
	}

}

func TestUpdateSummaryTrimsSpaces(t *testing.T) {

	uid := "x"
	summary := "x"

	task, _ := New(uid, summary)

	updatedSummary := strings.Repeat(" ", 50)

	UpdateSummary(&task, updatedSummary)

	if len(task.Summary) == len(summary) {
		t.Fatalf("Original summary string contains unneeded spaces, and should be trimmed, but output summary still contains those spaces")
	}

}
