package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	product "go-mico-shopping/product/proto/product"
	"go-micro-shopping/order/handler"
	order "go-micro-shopping/order/proto/order"
	"go-micro-shopping/order/repository"
	"log"
)

func main() {

	err := config.LoadFile("./config.json")
	if err != nil {
		log.Fatalf("Could not load config file: %s", err.Error())

	}
	conf := config.Map()
	db, err := createdatabase(conf["mysql"].(map[string]interface{}))
	if err != nil {
		log.Fatalf("can not find mysql: %s", err.Error())
	}
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.Version("latest"),
	)
	/**
	创建product的客户端，利用service,service既可以是变成服务，又可以变成客户端
	*/
	productCli := product.NewProductService("go.micro.srv.product", service.Client())

	order.RegisterOrderServiceHandler(service.Server(), &handler.Order{&repository.Order{db}, productCli})
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
func createdatabase(conf map[string]interface{}) (*gorm.DB, error) {
	host := conf["host"]
	port := conf["port"]
	user := conf["user"]
	dbName := conf["database"]
	password := conf["password"]
	return gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbName,
	),
	)
}
