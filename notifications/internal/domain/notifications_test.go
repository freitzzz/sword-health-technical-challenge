package domain

import (
	"testing"
)

func TestMarkAsReadCreatesNotificationReadStructWithUserIdAndNotificationId(t *testing.T) {
	tuid := "x"
	muid := "y"
	message := "x"

	notification := New(message, tuid)

	notificationRead := MarkAsRead(&notification, muid)

	if notificationRead.ManagerUserID != muid {
		t.Fatalf("Notification Read user id should be the same as the one marking it as read")
	}

}
