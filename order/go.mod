module go-micro-shopping/order

go 1.13

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	go-mico-shopping/product v0.0.0-00010101000000-000000000000
)

replace go-mico-shopping/product => ../product
