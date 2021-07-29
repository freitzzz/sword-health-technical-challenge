package amqp

import (
	"fmt"

	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/domain"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/logging"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func RegisterHandlers(c *amqp.Connection, db *gorm.DB) {

	ch, cerr := createChannel(c)

	if cerr != nil {
		logging.LogError("Failed to create channel")
		logging.LogError(cerr.Error())

		panic("AMQP channel is required to consume notifications")
	}

	nq, qerr := createNotificationsQueue(ch)

	if qerr != nil {
		logging.LogError("Failed to create notifications queue")
		logging.LogError(qerr.Error())

		panic("Notifications queue is required to consume notifications")
	}

	consumeNotifications(*ch, nq, db)

}

func consumeNotifications(ch amqp.Channel, q amqp.Queue, db *gorm.DB) chan bool {

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	if err != nil {
		logging.LogError("Failed to establish notifications channel")
		logging.LogError(err.Error())

		panic("Notifications channel connection is required to consume notifications")
	}

	forever := make(chan bool)

	go func() {

		defer close(forever)

		for d := range msgs {

			logging.LogInfo(fmt.Sprintf("Received message: %s", string(d.Body)))

			n, ferr := fromBytes(d.Body)

			if ferr != nil {
				logging.LogError("Failed to unmarshall notification message JSON")
			} else {
				_, ierr := data.InsertNotification(db, domain.New(n.Message, n.UserID))

				if ierr != nil {
					logging.LogError("Failed to insert notification in database")
				} else {
					logging.LogInfo("Inserted notification in database")

				}
			}

		}
	}()

	logging.LogInfo("Established notifications channel connection, waiting for messages")

	return forever

}
