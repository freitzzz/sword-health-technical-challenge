package amqp

import (
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/logging"
	"github.com/streadway/amqp"
)

func publishNotification(ch amqp.Channel, q amqp.Queue, n Notification) <-chan error {

	r := make(chan error)

	go func() {

		nb, merr := toBytes(n)

		if merr != nil {
			logging.LogError("Failed to marshal notification")
			logging.LogError(merr.Error())

			r <- merr
		} else {

			perr := ch.Publish("", q.Name, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        nb,
			})

			r <- perr

			if perr != nil {
				logging.LogError("Failed to publish notification")
				logging.LogError(perr.Error())

			}

		}

	}()

	return r

}
