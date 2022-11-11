
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

## gin的基本知识：

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

### p4 响应的多种格式

1. 三种数据类型的比较
```
map和python中字典：
    - map[key_type]value_type
struct结构体：
    - Struct {Attr1: value, Attr2: value}
python中属性：
    - person.age

属性必须是可用的标识符,不能是纯数字
键key只要是可哈希对象都可以,可以是纯数字
```



### p5 请求路由组的使用


### p6 中间件的编写与使用

1. 函数本身也可以作为类型, 用于声明变量和类型注解
```go
// github.com/gin-gonic/gin@v1.8.1/gin.go +43
// HandlerFunc defines the handler used by gin middleware as return value.
// HandlerFunc是接受Context指针的一类函数
type HandlerFunc func(*Context) 


// 函数RequestInfo返回值类型为HandlerFunc的函数
func RequestInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
        ......
	}
}

```


### p7

1. 下载go的mysql驱动包
```sh
go get "github.com/go-sql-driver/mysql"
```



## CloudRestaurant项目实战

### p8

1. 指针问题：
```go
package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

var _config *Config

func ParseConfig(file_path string) (*Config, error) {

	// 读取文件
	file, err := os.Open(file_path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	// decoder.Decode(_config) 会报错 panic: json: Unmarshal(nil *tool.Config)  TODO：指针的用法？
	if err := decoder.Decode(&_config); err != nil {
		return nil, err
	}
	fmt.Println(_config)
	return _config, nil
}


// marshal: v.结集；收集；安排；控制人群 n.空军元帅；司仪；典礼官；（美国法院的）执行官

```

1. new关键字
```go
func registerRouter(engine *gin.Engine) {
	new(controller.UserController).UserRouter(engine.Group("/user"))
}

在golang中提供了两种分配原语，即内建函数new 和 make(暂不提)。

- new是用来分配内存的内建函数，区别于其他语言中new会初始化内存，golang中的new只会将内存置零。

- 就是说，new(T)会为类型为T的新项分配已置零的内存空间，并返回它的地址。也就是一个类型为*T的值。

- 在GO中就是它返回一个指针，该指针指向新分配的，类型为T的零值。

零值属性带来的好处就是当我们用new声明某种类型时就已经分配好内存了，无需进一步处理即可正常工作。
```

### p10


1. 结构体的匿名字段(没有字段名，只有父级结构体类型)实现继承
```go
// 定义配置文件结构及参数
type Config struct {
	AppName    string       `json:"app_name"`
	AppMode    string       `json:"app_mode"`
	AppHost    string       `json:"app_host"`
	AppPort    string       `json:"app_port"`
	SMSCionfig `json:"sms"` //结构体的匿名字段实现继承
}

// 定义sms参数
type SMSCionfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	RegionID     string `json:"region_id"`
}

```


2. MVC的架构
```doc
前后端分析盛行之下：（view不在后端实现了）
- 中间件middleware：通用操作
- 路由router： 路由转发
- 控制器controller：处理路由
- 服务service：多（对象）表操作
- 数据对象dao：单（对象）表操作
- 模型model：表结构与结构体映射
```


3. 标准库
```go
json

rand

log

fmt

```


### p11


1. 返回值为nil时类型注解问题
```go
type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

func OrmEngine(conf Config) (Orm, error) {
	......
	return nil, err
}

// tool/OrmEngine.go:24:10: cannot use nil as Orm value in return statement
// 改正如下：修改返回值的类型
type Orm struct {
	*xorm.Engine
}

var DbEngine *Orm

func OrmEngine(conf Config) (*Orm, error) {
	......
	return nil, err
}

```

2. 结构体匿名字段初始化问题
```go
type MemberDao struct {
	*tool.Orm
}


member_dao := dao.MemberDao{tool.DbEngine}
result := member_dao.InsertCode(sms_cocde)

```

### p13 

1. 几种返回值的使用与接受处理
```go

nil

error

bool


```

2. 类型问题
```go
member := member_dao.QueryByPhhone(sms.Phone)  // member 这里由于QueryByPhhone返回值定义了其type为 *model.Member 后面再赋值时需要类型一致
if member.Id != 0 {
	return member
}

member = model.Member{}  // cannot use model.Member{} (value of type model.Member) as type *model.Member in assignmentgo
member.UserName = phone
member.Mobile = phone
member.RegisterTime = time.Now().Unix()


```

3. 当结构体作为参数传入函数，想用结构体存储处理后的数据时，需要使用结构体指针
```go
// 解析json格式参数
body := context.Request.Body
err := tool.ParseJson{}.Decode(body, smsLoginParam)   // smsLoginParam
if err != nil {
	tool.Failure(context, err.Error())
	return
}


// 解析json格式参数
body := context.Request.Body
err := tool.ParseJson{}.Decode(body, &smsLoginParam)  // 结构体是值类型，只有传递指针时，函数内修改了结构体的属性，才能保存下来
if err != nil {
	tool.Failure(context, err.Error())
	return
}


// smsLoginParam: CloudRestaurant/param.SmsLoginParam {Phone: "1234564564", Code: "510660"}
特别需要注意的点是：是否需要改变结构体属性并且保存下来，决定了是否必须使用指针
```

### p14 跨域

1. 相同域定义及跨域嗅探
```
协议、主机、端口三者都一致，定义为同域：
- http://www.baidu.com:80
- https://www.baidu.com:433
- wx://www.baidu.com:8090

TODO: github.com/gin-contrib/cors 源码
```





