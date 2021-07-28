package amqp

import "github.com/streadway/amqp"

const (
	notificationsQueue = "/notifications"
)

func createNotificationsQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(notificationsQueue, false, false, false, true, nil)
}
