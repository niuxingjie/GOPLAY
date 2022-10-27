
# Gin教程


- 安装与Demo
```sh
cd /mnt/d/ProgramData/GoPlay/qianphone/gin

mkdir GinHelloWorldCode

cd GinHelloWorldCode

go mod init GinHelloWorldCode

export GOPROXY="https://goproxy.cn,direct"

go get -u github.com/gin-gonic/gin

vim main.go

go mod tidy

git run main.go
```


### p2 HTTP请求和参数解析

1. package问题
- 案例1
```
// 一个文件夹下，多个文件可以属于同一个package
// gin文件夹下，多个文件都是"package gin"
// 同一个package内的变量可以相互引用而无需import

github.com/gin-gonic/gin.go 文件里用到了debugPrintWARNINGDefault
debugPrintWARNINGDefault 来自于github.com/gin-gonic/debug.go

```
- 疑问
```
github.com/gin-gonic/gins.go 文件夹下的Handle， GET， POST等函数是怎样绑定到engine.Handle上的？
```


2. 实现给获取不到的参数设置设置默认值
```go
// 获取不到值或返回错误，如果捕获错误则返回默认值 (等同于python中的异常处理)
// github.com/gin-gonic/context.go +422
func (c *Context) DefaultQuery(key, defaultValue string) string {
	if value, ok := c.GetQuery(key); ok {
		return value
	}
	return defaultValue
}
```

### p3 请求数据绑定与多数据格式处理

1. json数据类型和struct数据类型需一致
```go
{
    "name": "liming",
    "phone": 1708956235689,
    "password": "pass123456"
}
type Register struct {
	Username string `form:"name"` // 这里可以把name映射到Username
	Phone    int `form:"phone"`
	Password string `form:"password"`
}

二者不一样
{
    "name": "liming",
    "phone": "1708956235689",
    "password": "pass123456"
}
type Register struct {
	Username string `form:"name"` // 这里可以把name映射到Username
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

否则报错：
    2022/10/27 14:23:34 json: cannot unmarshal number into Go struct field Register.Phone of type string
    exit status 1
```










