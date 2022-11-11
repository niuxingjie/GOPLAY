
## START


## 03｜配好环境：选择一种最适合你的Go安装方法


###  配置 GOPROXY 这个 Go 环境变量

- ubuntu
```sh
export GOPROXY="socks5://127.0.0.1:1080,direct"


# go: finding module for package github.com/sirupsen/logrus
# github.com/bigwhite/module-mode imports
#         github.com/sirupsen/logrus: invalid proxy URL scheme (must be https, http, file): socks5://127.0.0.1:1080
unset GOPROXY

export GOPROXY="https://proxy.golang.org,direct"
unset GOPROXY


export GOPROXY="https://goproxy.cn,direct"

如下命令都会使用此代理：
go mod tidy
go install -v github.com/ramya-rao-a/go-outline@latest
```




## 04｜初窥门径：一个Go程序的结构是怎样的？



## 06｜构建模式：Go是怎么解决包依赖管理问题的？


### 创建你的第一个 Go Module


1. 新建项目及main.go文件
```go
// chapter06/main.go
package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)


func main() {
	logrus.Println("Hello ,go module mode\n")
	logrus.Println(uuid.NewString())
}

```

2. 通过 go mod init 创建 go.mod 文件，将当前项目变为一个 Go Module(TODO:项目和Module啥意思？)
```go
// 执行命令: go mod init github.com/bigwhite/module-mode
// go.mod
module github.com/bigwhite/module-mode  // 声明 module 路径 (module path)

go 1.18

```

3. go mod tidy 扫描源码，自动添加依赖到go.mod文件中,而且删除没有被使用的包

-  配置代理
```shell
export GOPROXY="https://goproxy.cn,direct"
go mod tidy
#    go: finding module for package github.com/sirupsen/logrus
#    go: downloading github.com/sirupsen/logrus v1.8.1
#    go: found github.com/sirupsen/logrus in github.com/sirupsen/logrus v1.8.1
#    go: downloading golang.org/x/sys v0.0.0-20191026070338-33540a1f6037
#    go: downloading github.com/stretchr/testify v1.2.2
#    go: downloading github.com/davecgh/go-spew v1.1.1
#    go: downloading github.com/pmezard/go-difflib v1.0.0
```

- 更新了go.mod
```go
// go.mod
module github.com/bigwhite/module-mode

go 1.18

require github.com/sirupsen/logrus v1.8.1

require golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 // indirect  TODO：对照日志这里为啥只有一个？
```

- 新增了go.sum
```go
// go.sum
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/sirupsen/logrus v1.8.1 h1:dJKuHgqk1NNQlqoA6BTlM1Wf9DOH3NBjQyu0h9+AZZE=
github.com/sirupsen/logrus v1.8.1/go.mod h1:yWOB1SBYBC5VeMP7gHvWumXLIWorT60ONWic61uBYv0=
github.com/stretchr/testify v1.2.2 h1:bSDNvY7ZPG5RlJ8otE/7V6gMiyenm9RtJ7IUVIAoJ1w=
github.com/stretchr/testify v1.2.2/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
golang.org/x/sys v0.0.0-20191026070338-33540a1f6037 h1:YyJpGZS1sBuBCzLAR1VEpK193GlqGZbnPFnPV/5Rsb4=
golang.org/x/sys v0.0.0-20191026070338-33540a1f6037/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
```

4. go build
```sh
go run main.go
# INFO[0000] Hello world!   
go build
ll
# total 1528
# drwxrwxrwx 1 root root     512 Jul 16 22:40 ./
# drwxrwxrwx 1 root root     512 Jul 16 22:01 ../
# -rwxrwxrwx 1 root root     164 Jul 16 22:31 go.mod*
# -rwxrwxrwx 1 root root     899 Jul 16 22:31 go.sum*
# -rwxrwxrwx 1 root root      99 Jul 16 22:37 main.go*
# -rwxrwxrwx 1 root root 2087136 Jul 16 22:40 module-mode*  # TODO：编译出来的可执行文件名取决于哪儿？

./module-mode  # TODO:文件名与go.mod里module path最后文件文一致？
# INFO[0000] Hello world!   
```

### 深入 Go Module 构建模式


### 特殊情况：使用 vendor

### QA

- 空导入问题
```text
Q1:空导入的方式的作用吗？我看很多源码中有使用这种包导入的方式。
A1: 
	像下面代码这样的包导入方式被称为“空导入”：
					
	import _ "foo"                                                                               
																						
	空导入也是导入，意味着我们将依赖foo这个路径下的包。但由于是空导入，我们并没有显式使用这个包中的任何语法元素。那么空导入的意义是什么呢？由于依赖foo包，程序初始化的时候会沿着包的依赖链初始化foo包，我们在08里会讲到包的初始化会按照常量->变量->init函数的次序进行。通常实践中空导入意味着期望依赖包的init函数得到执行，这个init函数中有我们需要的逻辑。
```

- vendor问题
```text
Q2: 在go module构建模式下，怎么对vendor目录的有无进行取舍呢？老师有什么实战建议呢？
A2：
	通常我们直接使用go module(非vendor)模式即可满足大部分需求。如果是那种开发环境受限，因无法访问外部代理而无法通过go命令自动解决依赖和下载依赖的环境下，我们通过vendor来辅助解决。
```

## 07｜构建模式：Go Module的6类常规操作


### 为当前 module 添加一个依赖

- code
```go
package main

import (

)

```




## 08｜入口函数与包初始化：搞清Go程序的执行次序


###  main.main 函数：Go 应用的入口函数

- Go 语言中有一个特殊的函数：main 包中的 main 函数，也就是 main.main，它是所有 Go 可执行程序的用户层执行逻辑的入口函数。
- 不过对于 main 包的 main 函数来说，你还需要明确一点，就是它虽然是用户层逻辑的入口函数，但它却不一定是用户层第一个被执行的函数。

- TODO：
    - "用户层"什么意思？
    - 用户层逻辑？
    - 用户层执行？


### init 函数：Go 包的初始化函数

- init基本介绍
```text
- 如果 main 包依赖的包中定义了 init 函数，或者是 main 包自身定义了 init 函数，
- 那么 Go 程序在这个包初始化的时候，就会自动调用它的 init 函数，
- 因此这些 init 函数的执行就都会发生在 main 函数之前。

- 不能手工显式地调用 init
```

- code
```go
package main

import "fmt"


func main() {
	fmt.Println("002-exec main()!")
}

func init() {
	// 初始化包：当我们要在 main.main 函数执行之前，执行一些函数或语句的时候，我们只需要将它放入 init 函数中就可以了。
	fmt.Println("001-init package!")
}
```

- 多包的init
```text
在初始化 Go 包时，Go 会按照一定的次序，逐一、顺序地调用这个包的 init 函数。
一般来说，先传递给 Go 编译器的源文件中的 init 函数，会先被执行；而同一个源文件中的多个 init 函数，会按声明顺序依次执行。

TODO: go源码是按照什么顺序传递给编译器编译的？
```

### Go 包的初始化次序

- ![Go 包的初始化次序](geek/img/e4ddb702876f4f2a0880e4353a390d0b.webp)

- 错误1
```text
main.go:6:2: import "github.com/bigwhite/prog-init-order/pkg1" is a program, not an importable package
pkg1.go中出现了“package main”的原因。一个main包，就是一个program
```

- code
```go
package main

import (
	"fmt"

	_ "github.com/bigwhite/prog-init-order/pkg1"
	_ "github.com/bigwhite/prog-init-order/pkg2"
)

var (
	_ = condtInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func condtInitCheck() string {
	if c1 != "" {
		// c1不是空值，所以说明已经初始化了 TODO：c1的空值为啥是是空字符串？
		fmt.Println("main: const c1 has been initialized")
	} 
	if c1 != "" {
		// c1不是空值，所以说明已经初始化了 TODO：c1的空值为啥是是空字符串？
		fmt.Println("main: const c1 has been initialized")
	} 
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s has been initialized \n", name)
	return name
}

func init() {
	fmt.Println("main: 001-first init func invoked")
}


func main() {
	fmt.Println("main: 002-main func invoked")
}
```

- 结果
```text
root@DESKTOP-BGUNPQG:# go run main.go 

	# main中首先是import导入包及包的初始化
	# 包内部，首先仍是是import导入包及包的初始化
	# 因此，深度遍历，最先执行的是pkg3

	main: pkg3 const c1 has been initialized  # 包内部，首先是包级别变量初始化
	main: pkg3 const c1 has been initialized
	main: var v1 has been initialized 
	main: var v2 has been initialized 
	main: pkg3 init func invoked              # 包内部，其次是init函数的的执行。（包内部main不会被执行）

	main: pkg2 const c1 has been initialized
	main: pkg2 const c1 has been initialized
	main: var v1 has been initialized 
	main: var v2 has been initialized 
	main: pkg2 init func invoked

	main: pkg1 const c1 has been initialized
	main: pkg1 const c1 has been initialized
	main: var v1 has been initialized 
	main: var v2 has been initialized 
	main: pkg1 init func invoked

	main: const c1 has been initialized
	main: const c1 has been initialized
	main: var v1 has been initialized 
	main: var v2 has been initialized 
	main: 001-first init func invoked
	main: 002-main func invoked               # main包内部的main才会被执行
```


### init 函数的三种用途

1. 重置包级变量值
```text
- 为什么需要重置：

- 如何重置：
	- 隐藏知识：包内部首先是变量初始化，其次是init函数执行，因此，init执行的时候，可以修改变量，从而实现重置包级变量值

```


2. 是实现对包级变量的复杂初始化

- 复杂初始化？
```text
- 复杂在哪里？
- 需要大量计算，才能得出变量初始化的值？
```

- 根据环境变量来决定初始化的值
```go
var (
    http2VerboseLogs    bool // 初始化时默认值为false
    http2logFrameWrites bool // 初始化时默认值为false
    http2logFrameReads  bool // 初始化时默认值为false
    http2inTests        bool // 初始化时默认值为false
)

func init() {
    e := os.Getenv("GODEBUG")
    if strings.Contains(e, "http2debug=1") {
        http2VerboseLogs = true // 在init中对http2VerboseLogs的值进行重置
    }
    if strings.Contains(e, "http2debug=2") {
        http2VerboseLogs = true // 在init中对http2VerboseLogs的值进行重置
        http2logFrameWrites = true // 在init中对http2logFrameWrites的值进行重置
        http2logFrameReads = true // 在init中对http2logFrameReads的值进行重置
    }
}
```

3. 在 init 函数中实现“注册模式”。
- 注册模式
```
- 设计模式？

- 注册是什么意? 

- 疑问：
	- 下面这个示例程序image支持 png、jpeg、gif 三种格式的图片，而达成这一目标的原因，正是 image/png、image/jpeg 和 image/gif 包都在各自的 init 函数中，将自己“注册”到 image 的支持格式列表中了
```

- code-main.go
```go

package main

import (
    "fmt"
    "image"
    _ "image/gif" // 以空导入方式注入gif图片格式驱动
    _ "image/jpeg" // 以空导入方式注入jpeg图片格式驱动
    _ "image/png" // 以空导入方式注入png图片格式驱动
    "os"
)

func main() {
    // 支持png, jpeg, gif
    width, height, err := imageSize(os.Args[1]) // 获取传入的图片文件的宽与高
    if err != nil {
        fmt.Println("get image size error:", err)
        return
    }
    fmt.Printf("image size: [%d, %d]\n", width, height)
}

func imageSize(imageFile string) (int, int, error) {
    f, _ := os.Open(imageFile) // 打开图文文件
    defer f.Close()

    img, _, err := image.Decode(f) // 对文件进行解码，得到图片实例
    if err != nil {
        return 0, 0, err
    }

    b := img.Bounds() // 返回图片区域
    return b.Max.X, b.Max.Y, nil
}
```

- code-空导入的依赖包
```go
// TODO: 整个项目中，包级变量（如：image是包名，还是包级变量？）是全局变量？全局变量可以在任何位置修改？


// $GOROOT/src/image/png/reader.go
func init() {
    image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}

// $GOROOT/src/image/jpeg/reader.go
func init() {
    image.RegisterFormat("jpeg", "\xff\xd8", Decode, DecodeConfig)
}

// $GOROOT/src/image/gif/reader.go
func init() {
    image.RegisterFormat("gif", "GIF8?a", Decode, DecodeConfig)
}  

```


## 09｜即学即练：构建一个Web服务就是这么简单


### 一个基于 HTTP 协议的 Web 服务

- 初始化项目
```sh

$ mkdir simple-http-server

$ cd simple-http-server

$ go mod init simple-http-server   # TODO：可以直接创建一个go module？
$ cat go.mod 
	module simple-http-server

	go 1.18

$ touch main.go

$ go run main.go

$ curl -I 127.0.0.1:8000
	HTTP/1.1 200 OK
	Date: Tue, 19 Jul 2022 13:33:37 GMT
	Content-Length: 12
	Content-Type: text/plain; charset=utf-8

$ go build

$ ./simple-http-server

$ curl 127.0.0.1:8000
	Hello Wordl!
```

- code
```go
package main

import "net/http"

func main()  {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello Wordl!"))  
	})
	http.ListenAndServe(":8000", nil)
}

```

- point:字节与字符串转换
```go
// [Golang中[]byte与string转换全解析](https://zhuanlan.zhihu.com/p/270626496)
// []byte 字符串转换为字节，类似python中的"string".encode("utf-8")

// string to []byte
s1 := "hello"
b := []byte(s1)

// []byte to string
s2 := string(b)
```


### 图书管理 API 服务

- 项目目录结构
```sh
tree ./ -L 3
./
├── cmd
│   └── bookstore        // 放置bookstore main包源码
│       └── main.go
├── go.mod
├── internal             // 存放项目内部包的目录
│   └── store
│       └── memstore.go
├── server               // HTTP服务器模块
│   ├── middleware
│   │   └── middleware.go
│   └── server.go
└── store                // 图书数据存储模块
    ├── factory
    │   └── factory.go
    └── store.go

8 directories, 7 files
```

- main包结构

![main包结构](./geek/img/5e8ee50b67a4229210b12afb94f55a19.webp)

- 图书数据存储模块（store)


- 知识点：Golang 单引号、双引号和反引号
```text
- 单引号(Single quote)，表示byte类型或rune类型，对应 uint8和int32类型，默认是 rune 类型。
	- byte用来强调数据是raw data，而不是数字；
	- 而rune用来表示Unicode的code point。
- 双引号(Double quote)，才是字符串，实际上是字符数组。可以用索引号访问某字节，也可以用len()函数来获取字符串所占的字节长度。
- 反引号(Back quote)，表示字符串字面量，但不支持任何转义序列。字面量 raw literal string 的意思是，你定义时写的啥样，它就啥样，你有换行，它就换行。你写转义字符，它也就展示转义字符。
```

- 知识点：面向接口编程
```text


```

- 知识点：工厂模式
```text



```

## 10｜变量声明：静态语言有别于动态语言的重要特征

- 静态语言有别于动态语言的一个重要特征：变量声明。

### 变量声明方法

0. 通用的变量声明方法

1. 省略类型信息的声明

2. 短变量声明


### 两类变量

0. 包级变量 (package varible)

1. 局部变量 (local varible)



## 11｜代码块与作用域：如何保证变量不会被遮蔽？


## 17 结构体

1. 结构体与接口的类型问题
```
判断变量的数据类型
语法：
1、值.(数据类型)，返回是否是指定的值的本身和数据类型
2、值.(type)，返回的是指的数据类型


Go 的类型转换常常让人有点迷，有两种形式的"类型转换"：
Type(obj) ：这种形式的类型转换要求 obj 对象的类型和 Type 是等价类型，即实现了相同的方法
obj.(Type) ：这种形式用于向下转型，即接口对象转结构体对象，所以 obj 必须是一个接口对象 , 这种形式在 Go 中一般叫做类型断言。
```


## 24 理解方法的本质

1. 方法本质上就是函数

```
- Go 函数的参数采用的是值拷贝传递

```


2. 选择 receiver 参数类型的第一个原则

```text
type T struct{}

func (t T) Tmethod {} 还是 func (t *T) Tmethod {}

- 1、如果 Go 方法要把对 receiver 参数代表的类型实例的修改，反映到原类型实例上，那么我们应该选择 *T 作为 receiver 参数的类型。
	- 无论是 T 类型实例，还是 *T 类型实例，都既可以调用 receiver 为 T 类型的方法，也可以调用 receiver 为 *T 类型的方法。
- 2、如果 receiver 参数类型的 size 较大
	- 因为函数是值拷贝传参，会增加参数值大小的内存
- 3、T 类型是否需要实现某个接口
	- 如果 T 类型需要实现某个接口，那我们就要使用 T 作为 receiver 参数的类型，来满足接口类型方法集合中的所有方法。
	- 如果 T 不需要实现某一接口，但 *T 需要实现该接口，那么根据方法集合概念，*T 的方法集合是包含 T 的方法集合的，这样我们在确定 Go 方法的 receiver 的类型时，参考原则一和原则二就可以了。

```


3. 方法集合
```go
// 可以获取【非接口类型】的方法集合的函数
func dumpMethodSet(i interface{}) {
    dynTyp := reflect.TypeOf(i)

    if dynTyp == nil {
        fmt.Printf("there is no dynamic type\n")
        return
    }

    n := dynTyp.NumMethod()
    if n == 0 {
        fmt.Printf("%s's method set is empty!\n", dynTyp)
        return
    }

    fmt.Printf("%s's method set:\n", dynTyp)
    for j := 0; j < n; j++ {
        fmt.Println("-", dynTyp.Method(j).Name)
    }
    fmt.Printf("\n")
}


```


## 29｜接口：为什么nil接口不等于nil？


1. 接口的【静态特性】与【动态特性】

- 静态特性
	```
	编译阶段的类型检查：
		- var err error  变量定义
		- func (err error) {} 函数参数校验
		- type Response struct { Err error} 结构体属性校验
		。。。。。。

	【接口interface】拥有静态类型：
		- 那就意味着编译器会在编译阶段对所有接口类型变量的赋值操作进行类型检查，
		- 编译器会检查右值的类型是否实现了该接口方法集合中的所有方法。
	
	空接口interface}{}可以持有任何类型数据
	```

- code
	```go
	package main

	import (
		"fmt"
	)

	func main() {
		fmt.Println("结构体的类型判断：")
		liming := Person{Name: "liming", Age: 28}
		Skill(liming)
	}

	type PersonAction interface {
		Run()
		Eat()
	}

	type Person struct {
		Name string
		Age  int
	}

	func (p Person) Run() {
		fmt.Println(p.Name, " can run")
	}

	func (p Person) Eat() {
		fmt.Println(p.Name, " can eat!")
	}

	// 接口的静态特性：PersonAction持有实现了接口方法的结构体
	// 编写Skil(Person{Name: "liming", Age: 20})是可以通过编译的
	func Skill(p PersonAction) {
		p.Run()
		p.Eat()
	}

	```

- 动态特性
  ```
	【接口】动态特性：
		- 就体现在接口类型变量在运行时还存储了右值的真实类型信息，这个右值的真实类型被称为接口类型变量的动态类型

  ```


## END





