module go-micro-shopping/notification

go 1.13

require (
	github.com/LiRonaldo/l-log v1.1.2
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	product v0.0.0-00010101000000-000000000000
	user v0.0.0-00010101000000-000000000000
)

replace (
	product => ../product
	user => ../user
)
