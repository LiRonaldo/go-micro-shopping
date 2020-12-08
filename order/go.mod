module go-micro-shopping/order

go 1.13

require (
	contrib.go.opencensus.io/exporter/zipkin v0.1.2
	github.com/LiRonaldo/l-log v1.1.2
	github.com/bwmarrin/snowflake v0.3.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/openzipkin/zipkin-go v0.2.2
	go-mico-shopping/product v0.0.0-00010101000000-000000000000
	go.opencensus.io v0.22.4

)

replace go-mico-shopping/product => ../product
