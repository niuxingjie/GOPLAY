
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
```

- 配置git仓库
```shell
git init

```






















