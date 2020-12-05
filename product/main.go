package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"go-micro-shopping/product/handler"
	"go-micro-shopping/product/subscriber"

	product "go-micro-shopping/product/proto/product"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.product"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), new(handler.Product))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.product", service.Server(), new(subscriber.Product))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.product", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
