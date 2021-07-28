package amqp

import "encoding/json"

type Notification struct {
	UserID  string `json:"userId"`
	Message string `json:"message"`
}

func fromBytes(nb []byte) (Notification, error) {
	notification := Notification{}
	uerr := json.Unmarshal(nb, &notification)

	return notification, uerr
}
