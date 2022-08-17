
## Beego教程

[beego 官方网址](https://beego.vip/)

注意：视频教程里的go/beego/bee的版本都较低，代码需要适当调整

- go path
```sh
tree /root/go/ -L 3
    /root/go/
    ├── bin
    │   ├── go-outline
    │   ├── gocode
    │   ├── godef
    │   └── goimports
    └── pkg
        ├── mod
        │   ├── 9fans.net
        │   ├── cache
        │   ├── github.com
        │   ├── go.starlark.net@v0.0.0-20190702223751-32f345186213
        │   ├── golang.org
        │   ├── google.golang.org
        │   └── gopkg.in
        └── sumdb
            └── sum.golang.org

    12 directories, 4 files

# 将 /root/go文件夹添加到当前工作区，方便查看go包的源码
```


- 安装beego
```shell
go get github.com/beego/beego/v2@v2.0.0

    go: go.mod file not found in current directory or any parent directory.
            'go get' is no longer supported outside a module.
            To build and install a command, use 'go install' with a version,
            like 'go install example.com/cmd@latest'
            For more information, see https://golang.org/doc/go-get-install-deprecation
            or run 'go help get' or 'go help install'.

go version
    go version go1.18.3 linux/amd64

mkdir -p beego/FirstBeego

cd beego/FirstBeego

go mod init FirstBeego
    go: creating new go.mod: module FirstBeego


export GOPROXY="https://goproxy.cn,direct"
go get github.com/beego/beego/v2@v2.0.0

tree ./
    ./
    ├── go.mod
    └── go.sum

```

- 编写并运行第一个web
```shell
go run main.go

/root/go/pkg/mod/github.com/beego/beego/v2@v2.0.0/server/web/context/output.go:34:2: missing go.sum entry for module providing package gopkg.in/yaml.v2 (imported by github.com/beego/beego/v2/server/web/context); to add:
        go get github.com/beego/beego/v2/server/web/context@v2.0.0

go mod tidy

go run main.go
```


- 安装bee开发工具
```shell

export GOPROXY="https://goproxy.cn,direct"
go get github.com/beego/bee  # 这个命令不会安装bee命令到/root/go/bin下

go install github.com/beego/bee
    go: 'go install' requires a version when current directory is not in a module
            Try 'go install github.com/beego/bee@latest' to install the latest version
            
go install github.com/beego/bee@v1.12.3  #

/root/go/bin/bee -h

/root/go/bin/bee version

cd /mnt/d/ProgramData/GoPlay/qianphone/beego

/root/go/bin/bee new SecondBeego

tree SecondBeego/ -L 2
    SecondBeego/
    ├── conf
    │   └── app.conf
    ├── controllers
    │   └── default.go
    ├── go.mod
    ├── main.go
    ├── models
    ├── routers
    │   └── router.go
    ├── static
    │   ├── css
    │   ├── img
    │   └── js
    ├── tests
    │   └── default_test.go
    └── views
        └── index.tpl

    10 directories, 7 files

cd SecondBeego

/root/go/bin/bee run
    2022/08/16 21:55:03 ERROR    ▶ 0001 GOPATH environment is empty,may be you use `go module`
    ______
    | ___ \
    | |_/ /  ___   ___
    | ___ \ / _ \ / _ \
    | |_/ /|  __/|  __/
    \____/  \___| \___| v1.12.0
    2022/08/16 21:55:03 WARN     ▶ 0002 Running application outside of GOPATH
    2022/08/16 21:55:03 INFO     ▶ 0003 Using 'SecondBeego' as 'appname'
    2022/08/16 21:55:03 INFO     ▶ 0004 Initializing watcher...
    controllers/default.go:4:2: missing go.sum entry for module providing package github.com/astaxie/beego (imported by SecondBeego); to add:
            go get SecondBeego
    2022/08/16 21:55:04 ERROR    ▶ 0005 Failed to build the application: controllers/default.go:4:2: missing go.sum entry for module providing package github.com/astaxie/beego (imported by SecondBeego); to add:
            go get SecondBeego

go get SecondBeego

/root/go/bin/bee run
    2022/08/16 21:56:43 SUCCESS  ▶ 0005 Built Successfully!
    2022/08/16 21:56:43 INFO     ▶ 0006 Restarting 'SecondBeego'...
    2022/08/16 21:56:43 SUCCESS  ▶ 0007 './SecondBeego' is running...
    2022/08/16 21:56:43.393 [I] [asm_amd64.s:1571]  http server Running on http://:8080
    2022/08/16 21:57:22.723 [D] [server.go:2916]  |      127.0.0.1| 200 |     7.7627ms|   match| GET      /     r:/
    2022/08/16 21:57:22.745 [D] [server.go:2916]  |      127.0.0.1| 200 |        540µs|   match| GET      /static/js/reload.min.js

```


- MysqlDemo数据库链接
```sh

cd /mnt/d/ProgramData/GoPlay/qianphone/beego

/root/go/bin/bee new MysqlDemo

cd MysqlDemo/

export GOPROXY="https://goproxy.cn,direct"

go get MysqlDemo  # TODO：貌似这个命令会把本项目相关依赖加入mod中？
    go: added github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18

/root/go/bin/bee run


# 直接安装依赖会报错
go install github.com/go-sql-driver/mysql
missing go.sum entry for module providing package github.com/go-sql-driver/mysql; to add:
        go mod download github.com/go-sql-driver/mysql

go mod download github.com/go-sql-driver/mysql
```

- myblogweb开发实战
```



```



