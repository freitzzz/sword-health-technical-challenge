package amqp

import "github.com/streadway/amqp"

func createChannel(c *amqp.Connection) (*amqp.Channel, error) {

	return c.Channel()
}
