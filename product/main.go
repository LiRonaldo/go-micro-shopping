package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"go-micro-shopping/product/handler"
	"go-micro-shopping/product/model"
	proto "go-micro-shopping/product/proto/product"
	"go-micro-shopping/product/repository"
)

func main() {

	err := config.LoadFile("./config.json")
	if err != nil {
		log.Fatal("config file not find")
	}
	conMap := config.Map()
	db, err := createdatabase(conMap["mysql"].(map[string]interface{}))
	db.AutoMigrate(&model.Product{})
	defer db.Close()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.product"),
		micro.Version("latest"),
	)
	service.Init()
	product := &handler.Product{Pro: &repository.Product{db}}
	proto.RegisterProductServiceHandler(service.Server(), product)
	if err := service.Run(); err != nil {
		log.Fatal(err)
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
