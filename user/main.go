package main

import (
	"fmt"
	_ "github.com/LiRonaldo/l-log"
	log "github.com/LiRonaldo/l-log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"go-micro-shopping/user/handler"
	"go-micro-shopping/user/model"
	user "go-micro-shopping/user/proto/user"
	"go-micro-shopping/user/repository"
)

/**
根据proto定义好的方法，入参和出参，生成的文件，文件中包含服务端和客户端的生成方法。
服务端RegisterUserServiceHandler将服务，注册进去，只会返回err。

客户端NewUserService 创建，根据服务名获得，传入你newService的服务的client，生成一个子类，
然后调用定义好的方法。，客户端和服务端方法名一样就是proto定义的方法，只是入参返回不一样。
客户端是返回resp，服务端resp也是入参。


go run main.go 启动服务。
然后变成api 调用  micro api --namespace=go.micro.srv
或者 写一个客户端，调用，具体是  micro.NewService，返回一个service，然后调用proto文件里的newService方法。将service.client传入，
返回一个重写了proto文件里的client ，然后调用
或者直接将新建的client 变成api
*/

func main() {
	//加载配置文件
	err := config.LoadFile("./config.json")
	if err != nil {
		log.Fatal("Could not load config file: %s", err.Error())
	}
	conf := config.Map()

	db, err := createdatabase(conf["mysql"].(map[string]interface{}))

	defer db.Close()
	//自动创建表
	db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("connection error : %v \n", err)
	}
	repo := &repository.User{db}
	// New Service
	//consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	service.Init()
	user.RegisterUserServiceHandler(service.Server(), &handler.User{repo})
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
