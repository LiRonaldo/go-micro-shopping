package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"go-micro-shopping/notification/handler"
	"go-micro-shopping/notification/subscriber"

	notification "go-micro-shopping/notification/proto/notification"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.notification"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	notification.RegisterNotificationHandler(service.Server(), new(handler.Notification))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.notification", service.Server(), new(subscriber.Notification))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.notification", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
