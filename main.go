package main

import (
	"fmt"
	"os"

	"github.com/getzion/relay/api"
	"github.com/sirupsen/logrus"
)

func main() {
	server := api.InitDWNServer()
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	logrus.Fatal(server.Listen(fmt.Sprintf("%s:%s", host, port)))
}
