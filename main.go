package main

import (
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
	"github.com/tonymj76/micro-postgres/datastore"
	"github.com/tonymj76/micro-postgres/handler"
	pbUser "github.com/tonymj76/micro-postgres/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("service.user"),
		micro.Version("0.1"),
	)

	// Initialise service
	service.Init()

	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.StampMilli,
		FullTimestamp:   true,
	}
	conn, err := datastore.NewConnection(log)
	defer conn.Close()
	log.WithError(err).Fatal("database faild to connect")
	// Register Handler
	pbUser.RegisterUserServiceHandler(service.Server(), &handler.Service{R: conn})

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("service.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
