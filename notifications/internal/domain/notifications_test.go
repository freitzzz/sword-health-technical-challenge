package domain

import (
	"testing"
)

func TestMarkAsReadCreatesNotificationReadStructWithUserIdAndNotificationId(t *testing.T) {
	uid := "x"
	message := "x"

	notification := New(message)

	notificationRead := MarkAsRead(&notification, uid)

	if notificationRead.UserID != uid {
		t.Fatalf("Notification Read user id should be the same as the one marking it as read")
	}

}
