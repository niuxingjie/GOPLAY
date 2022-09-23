
### 《Go语言极简一本通》系列

[作者网站](http://www.go-edu.cn/about/)
[go_basic_edu](https://github.com/i-coder-robot/go_basic_edu)



### 初始化

```sh
go env

cd /workspaces/GOPLAY/go-edu

go mod init goedu

go mod tidy
```


### 记录


go标准库不在GOPATH下，而是在GOROOT下。
具体的位置：/usr/local/go/src

GOPAT文件夹有点类似虚拟环境的作用，把同一项目的依赖都放在此目录下。
公用的标准库，仍然放在GOROOT下。
这与go是编译型语言有关系吧，只要控制好编译时，包路径就好了。

