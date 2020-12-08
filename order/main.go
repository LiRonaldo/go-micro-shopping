package main

import (
	"contrib.go.opencensus.io/exporter/zipkin"
	"fmt"
	_ "github.com/LiRonaldo/l-log"
	log "github.com/LiRonaldo/l-log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-plugins/broker/mqtt"
	_ "github.com/micro/go-plugins/broker/mqtt"
	product "go-mico-shopping/product/proto/product"
	"go-micro-shopping/order/handler"
	"go-micro-shopping/order/model"
	order "go-micro-shopping/order/proto/order"
	"go-micro-shopping/order/repository"
	"go.opencensus.io/trace"
	//wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opencensus"
	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
	"os"
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
	db.AutoMigrate(&model.Order{})
	defer db.Close()
	// New Service
	mq := mqtt.NewBroker()

	service := micro.NewService(
		micro.Name("go.micro.srv.order"),
		micro.Version("latest"),
		micro.Broker(mq),
	)
	/**
	创建product的客户端，利用service,service既可以是变成服务，又可以变成客户端
	*/
	productCli := product.NewProductService("go.micro.srv.product", service.Client())
	/**
	消息发布者
	*/
	publisher := micro.NewPublisher("notification.submit", service.Client())
	/**
	  传给order
	*/
	order.RegisterOrderServiceHandler(service.Server(), &handler.Order{&repository.Order{db}, productCli, publisher})
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

func TraceBoot() {
	apiURL := "http://192.168.0.111:9411/api/v2/spans"
	hostPort, _ := os.Hostname()
	serviceName := "go.micro.srv.order"

	localEndpoint, err := openzipkin.NewEndpoint(serviceName, hostPort)
	if err != nil {
		log.Fatalf("Failed to create the local zipkinEndpoint: %v", err)
	}
	reporter := zipkinHTTP.NewReporter(apiURL)
	ze := zipkin.NewExporter(reporter, localEndpoint)
	trace.RegisterExporter(ze)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	return
}
