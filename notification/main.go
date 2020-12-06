package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/broker/mqtt"
	_ "github.com/micro/go-plugins/broker/mqtt"
	"go-micro-shopping/notification/subscriber"
	"log"
)

func main() {

	mq := mqtt.NewBroker()
	mq.Init()
	mq.Connect()
	service := micro.NewService(
		micro.Name("go.micro.srv.notification"),
		micro.Version("latest"),
		micro.Broker(mq),
	)
	service.Init()
	micro.RegisterSubscriber("notification.submit", service.Server(), &subscriber.Notification{})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
