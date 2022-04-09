package main

import (
	"fmt"
	"os"

	"github.com/getzion/relay/api/dwn/server"
	"github.com/sirupsen/logrus"
)

func main() {
	server := server.InitDWNServer()
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	logrus.Fatal(server.Listen(fmt.Sprintf("%s:%s", host, port)))
}
