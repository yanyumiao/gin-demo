gin-demo  
A demo about use gin to develop web service  

1 controller  
Multiple files under one directory  

2 router.go  
Separate router from main.go

3 gin.H  
H is a shortcut for map[string]interface{}  

4 包下多文件及其方法命名的思考  
如model有person.go 包含了 AddPerson DelPerson等方法  
user.go 包含 AddUser DelUser等方法    
方法和文件对应清晰 方便后期项目的维护  

5 async demo  