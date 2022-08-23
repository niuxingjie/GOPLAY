
### go 的安装

- 打开wsl终端
```shell
cd /mnt/d/ProgramData/GoPlay

ubuntu install go
sudo add-apt-repository ppa:longsleep/golang-backports 
sudo apt-get install golang-go
```

- shell 配置代理
```shell
export http_proxy='socks5://127.0.0.1:1080'
export https_proxy='socks5://127.0.0.1:1080'


unset http_proxy
unset https_proxy


# 上面的不太好用，下面的方式也很好
export GOPROXY="https://goproxy.cn,direct" 
# 如下命令都会使用此代理：
go mod tidy

go install -v github.com/ramya-rao-a/go-outline@latest
```

- vscode给go配置跳转代码提示等插件
```sh
# ubuntu wsl2
export GOPROXY="https://goproxy.cn,direct" 
go install -v github.com/ramya-rao-a/go-outline@latest
go install -v golang.org/x/tools/cmd/goimports@latest
go install -v github.com/rogpeppe/godef@latest
go install -v github.com/stamblerre/gocode@latest

插件特性：
    - ctrl + s 自动格式化源码中的go代码
```

- 配置git仓库
```shell
git init

git remote add origin https://github.com/niuxingjie/GOPLAY.git

git push -u origin master
```

- git 配置本地代理
```shell
git config --global http.proxy 'http://127.0.0.1:1080'
git config --global https.proxy 'https://127.0.0.1:1080'

git config --global http.proxy 'socks5://127.0.0.1:1080'
git config --global https.proxy 'socks5://127.0.0.1:1080'

git config --global --unset http.proxy
git config --global --unset https.proxy

```


- git 忽略文件
```shell
$ git add .
    On branch dev
    Changes to be committed:
    (use "git reset HEAD <file>..." to unstage)

            modified:   .gitignore
            new file:   .vscode/configurationCache.log
            new file:   .vscode/settings.json
            new file:   .vscode/targets.log
            modified:   README.md
            modified:   "Tony Bai \302\267 Go \350\257\255\350\250\200\347\254\254\344\270\200\350\257\276.md"
            new file:   geek/chapter09/bookstore/Makefile
            new file:   geek/chapter09/bookstore/bookstore

$ git status

$ vim gitignore
    .vscode/*

$ git rm -r --cached .vscode

$ git status
$ git status
    On branch dev
    Changes to be committed:
    (use "git reset HEAD <file>..." to unstage)

            modified:   .gitignore
            modified:   README.md
            modified:   "Tony Bai \302\267 Go \350\257\255\350\250\200\347\254\254\344\270\200\350\257\276.md"
```


















