gin-demo  
使用gin框架开发web application的demo  

1 controller  
controller目录下包含多个文件  
测试同一个目录(包)下,包含多个go文件但是包名一致的情况  

2 router.og  
通过将路由信息写入router.go, 从而将路由从main.go中分离出来  
这个方法可以借鉴,可以动手尝试使用这个方法分离httprouter的路由配置  

3 gin.H  
H is a shortcut for map[string]interface{}  

4 model下多文件及其方法命名的思考  
如当前model下只有person.go 包含了 AddPerson DelPerson等方法  
如果将来再添加user.go文件 那么user.go文件中方法命名 AddUser DelUser  
这样一来 虽然只有一个包 却能通过命名规则和文件对应起来 方便后期项目的维护和管理  
