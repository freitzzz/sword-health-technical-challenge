package http

import (
	"fmt"
	"os"
)

const (
	serverHostEnvKey = "server_host"
	serverPortEnvKey = "server_port"
)

var (
	serverHost = os.Getenv(serverHostEnvKey)
	serverPort = os.Getenv(serverPortEnvKey)
)

func ServerAddress() string {
	return fmt.Sprintf("%s:%s", serverHost, serverPort)
}
