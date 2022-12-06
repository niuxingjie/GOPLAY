

## 资料

[Bilibili视频]()


## p8 第一个hello world 三种import


## p9 go执行原理和命令

```sh
whereis go

tree /usr/lib/go-1.18/ -L 2
    /usr/lib/go-1.18/
    ├── VERSION
    ├── api -> ../../share/go-1.18/api
    ├── bin
    │   ├── go
    │   └── gofmt
    ├── misc -> ../../share/go-1.18/misc
    ├── pkg
    │   ├── include -> ../../../share/go-1.18/pkg/include
    │   ├── linux_amd64
    │   └── tool
    ├── src -> ../../share/go-1.18/src
    └── test -> ../../share/go-1.18/test

tree /usr/share/go-1.18/ -L 1
    /usr/share/go-1.18/
    ├── api
    ├── misc
    ├── pkg
    ├── src       # 包源码
    └── test


go env # 查看当前配置的go的环境变量信息
    GO111MODULE=""
    GOARCH="amd64"
    GOCACHE="/root/.cache/go-build"
    GOENV="/root/.config/go/env"
    GOHOSTOS="linux"
    GOMODCACHE="/root/go/pkg/mod"
    GOOS="linux"
    GOPATH="/root/go"
    GOPRIVATE=""
    GOPROXY="https://proxy.golang.org,direct"
    GOROOT="/usr/lib/go-1.18"
    GOSUMDB="sum.golang.org"
    GOTMPDIR=""
    GOTOOLDIR="/usr/lib/go-1.18/pkg/tool/linux_amd64"
    GOVCS=""
    GOVERSION="go1.18.3"
    GCCGO="gccgo"
    CGO_ENABLED="1"
    GOMOD="/dev/null"
    GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build24339221=/tmp/go-build -gno-record-gcc-switches"
    ......

go -h  # 查看更多命令信息

```


## p12 编码规范


## p16 iota关键字

## p18 string类型和字符编码

## p29 if的其他写法

## p30 swtich分支语句

## p31 swtich的其他写法

## p32 swtich的break和fallthrough

## P33 for循环语句

## p34 for 语句的其他写法

## p41 goto语句


## p42 生成随机数


## p43 Array的初步使用


## p45 数据的遍历for range


## p45 数组是值类型

- 值类型意味着赋值时是值传递（对应的是：引用传递，map,slice）
```go
package main

import "fmt"


func main(){
    var arr1 = [5]int{1,2,3,4,5}
    var arr2 = [5]int{6,7,8,9,10}

    fmt.Println(arr1)
    fmt.Println(arr2)

    arr2 = arr1  // 值传递
    fmt.Println(arr2)

    arr2[0] = 0
    fmt.Println(arr1)  // arr1没有变，说明值传递
    fmt.Println(arr2)
}

```


## p49 Slice初步使用

- 内建函数make创建引用类型数据
- 切片本身是不存储数据的，是切片底层的数组存储数据


## p88 方法


## p50 Slice的内存分析








## END