package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/sirupsen/logrus"
	"github.com/tonymj76/micro-postgres/datastore"
	"github.com/tonymj76/micro-postgres/handler"
	_ "github.com/tonymj76/micro-postgres/log"
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
	conn, err := datastore.NewConnection(log.New())
	if err != nil {
		log.WithError(err).Fatal("database failed to connect")
	}
	defer conn.Close()
	// Register Handler
	pbUser.RegisterUserServiceHandler(service.Server(), &handler.Service{R: conn})

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("service.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.WithError(err).Fatal("unable to run serive")
	}
}
