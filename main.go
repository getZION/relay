package main

import (
	"github.com/getzion/relay/api/dwn/server"
	"github.com/getzion/relay/api/models"
	"github.com/getzion/relay/api/storage"
	"github.com/sirupsen/logrus"
)

func main() {
	storage, err := storage.NewStorage("mysql")
	if err != nil {
		logrus.Panic(err)
	}

	modelManager := models.NewModelManager(storage)
	server := server.InitDWNServer(modelManager, storage)

	logrus.Fatal(server.Listen(":8080"))
}
