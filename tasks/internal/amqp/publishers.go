package amqp

import (
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/logging"
	"github.com/streadway/amqp"
)

func PublishNotification(mb MailBox, n Notification) chan error {

	r := make(chan error)

	go func() {

		defer close(r)

		nb, merr := toBytes(n)

		if merr != nil {
			logging.LogError("Failed to marshal notification")
			logging.LogError(merr.Error())

			r <- merr
		} else {

			perr := mb.Channel.Publish("", mb.Queue.Name, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        nb,
			})

			if perr != nil {

				logging.LogError("Failed to publish notification")
				logging.LogError(perr.Error())

			} else {

				logging.LogWarning("Published notification")
			}

		}

	}()

	return r

}
