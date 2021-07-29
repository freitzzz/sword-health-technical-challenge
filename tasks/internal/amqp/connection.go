package amqp

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

const (
	rabbitMQUsernameEnvKey = "rabbitmq_user"
	rabbitMQPasswordEnvKey = "rabbitmq_pass"
	rabbitMQAddressEnvKey  = "rabbitmq_address"
)

var (
	rabbitMQUsername = os.Getenv(rabbitMQUsernameEnvKey)
	rabbitMQPassword = os.Getenv(rabbitMQPasswordEnvKey)
	rabbitMQAddress  = os.Getenv(rabbitMQAddressEnvKey)
)

func OpenMQConnection() (*amqp.Connection, error) {

	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", rabbitMQUsername, rabbitMQPassword, rabbitMQAddress))
}
