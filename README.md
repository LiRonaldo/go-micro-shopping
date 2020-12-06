# go-micro-shopping
#<h3>go-micro 实战。
<h4>user服务
<h5>执行命令protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
<h5>读取配置文件:user/main.go
<h5>db.AutoMigrate(&model.User{}) 自动创建表 不过会戴上s ，变成users，time.time属性对应库里datetime，
repo.Dao.Model(user).Updates(&user) ，model传入是指针，update传入是内容值  &指针获取的是内容值
<h5>run mian.go  将服务启动， 
micro api     --namespace=go.micro.srv  
api调用 设置命名空间，
服务名是 micro.Name("go.micro.srv.user"), 
请求调用的时候可不用写命名空间那一部分，直接写user/服务名（注意大小写）/方法名（注意大小写)相当于省去go.micro.srv

<h4>product服务
<h2>终于碰到坑了
<h5>product对应表里的字段，因为引用gorm框架，所以product struct 引用 gorm.model,会自带id，
CreatedAt time.Time，UpdatedAt time.Time，DeletedAt *time.Time，使用AutoMigrate时会自己建表，并自己创建这三个字段。
proto文件中也有message product，用proto文件里的product，和model里的product 去掺入参数一样的，都能达到要求。
proto里的product 中多余的XX开头的字段不会显示到页面中，
<h4>但是如果使用gorm框架的方法时使用model里的product当传入参数时要注意：因为model里的product引入了gorm.model，所以查询的时候会附加一个 DeletedAt is null 的条件，要注意。
gorm 底层也是拼接的sql语句，之前写过一个类似的orm框架，走debug就会发现is null语句.
gorm框架的方法find，creat update where 方法要根据文档要求传参数带不带&，不然会报错
<h4>order服务
<h5>在order服务想引用product中的东西。直接improt是不会生效的。将order mod文件加入 replace go-mico-shopping/product => ../product
说白了就是将improt 导入的包替换成本地的
                                                                   
                                                                    