# go-micro-shopping
#<h3>go-micro 实战。
<h5>执行命令protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
<h5>读取配置文件:user/main.go
<h5>db.AutoMigrate(&model.User{}) 自动创建表 不过会戴上s ，变成users，time.time属性对应库里datetime，
repo.Dao.Model(user).Updates(&user) ，model传入是指针，update传入是内容值  &指针获取的是内容值
<h5>run mian.go  将服务启动， 
micro api     --namespace=go.micro.srv  
api调用 设置命名空间，
服务名是 micro.Name("go.micro.srv.user"), 
请求调用的时候可不用写命名空间那一部分，直接写user/服务名（注意大小写）/方法名（注意大小写)相当于省去go.micro.srv
