gin-demo
使用gin框架开发web application的demo

1 controller
controller目录下包含 hello.go和test.go两个文件
测试同一个目录(包)下,包含多个go文件但是包名一致的情况

2 router.og
通过将路由信息写入router.go, 从而将路由从main.go中分离出来
这个方法可以借鉴,可以动手尝试使用这个方法分离httprouter的路由配置 

3 gin.H 
H is a shortcut for map[string]interface{}
参考test.go的Pong方法 和 其源码实现
