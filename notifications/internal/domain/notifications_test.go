package domain

import (
	"testing"
)

func TestMarkAsReadMarksReadPropertyAsTrue(t *testing.T) {
	uid := "x"
	message := "x"

	notification, _ := New(uid, message)

	readPropBefore := notification.Read

	MarkAsRead(&notification)

	readPropAfter := notification.Read

	if readPropBefore == readPropAfter {
		t.Fatalf("Read properties before and after should be different")
	}

	if !readPropAfter {
		t.Fatalf("MarkAsRead procedure should mark Read property as false, but condition is not met")
	}

}
