package amqp

import (
	"encoding/json"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/logging"
	"github.com/streadway/amqp"
)

type Notification struct {
	UserID  string `json:"userId"`
	Message string `json:"message"`
}

type MailBox struct {
	Channel amqp.Channel
	Queue   amqp.Queue
}

func toBytes(n Notification) ([]byte, error) {

	return json.Marshal(n)

}

func SetupMailBox(c *amqp.Connection) (MailBox, error) {

	ch, cerr := createChannel(c)

	var mb MailBox

	if cerr != nil {
		logging.LogError("Failed to create channel")
		logging.LogError(cerr.Error())

		return mb, cerr

	}

	nq, qerr := createNotificationsQueue(ch)

	if qerr != nil {
		logging.LogError("Failed to create notifications queue")
		logging.LogError(qerr.Error())

		return mb, qerr

	}

	mb = MailBox{Channel: *ch, Queue: nq}

	return mb, nil

}
