#### gin-demo   
一个使用gin框架开发web服务的demo
* MVC分层
* 统一输出格式
* 支持MySQL Rdis
----
##### 开始
1 DB准备 
* demo使用MySQL数据库 
* 先创建 test数据库 运行 person.sql 创建person表

2 go 环境设置

项目进行了更新 现在采用了 go module 进行工程管理
所以启动项目之前 要设置支持 module

* 必要条件 go version >= v1.11
* 设置go 支持 module
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

2.3 注意

因为采用了 go module 管理工程 项目位置不要再放置在 $GOPATH/src

3 运行
因为采用了go module 所以直接在main.go所在目录
```
go run main.go
```
访问 127.0.0.1:10086 看到 hello 输出 成功!