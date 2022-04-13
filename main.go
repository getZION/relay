package main

import (
	"fmt"
	"os"

	"github.com/getzion/relay/api/dwn/server"
	"github.com/getzion/relay/api/models"
	"github.com/getzion/relay/api/storage"
	"github.com/getzion/relay/api/validator"
	"github.com/sirupsen/logrus"
)

func main() {
	validator.InitValidator()

	storage, err := storage.NewStorage("mysql")
	if err != nil {
		logrus.Panic(err)
	}

	modelManager := models.NewModelManager(storage)
	server := server.InitDWNServer(modelManager, storage)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	logrus.Fatal(server.Listen(fmt.Sprintf("%s:%s", host, port)))
}
