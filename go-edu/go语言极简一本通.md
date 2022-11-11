
### 《Go语言极简一本通》系列

[作者网站](http://www.go-edu.cn/about/)
[go_basic_edu](https://github.com/i-coder-robot/go_basic_edu)



### 初始化

- go
```sh
go env

cd /workspaces/GOPLAY/go-edu

go mod init goedu

go mod tidy
```

- 查网络
```sh
cat /etc/hosts

    127.0.0.1       localhost
    ::1     localhost ip6-localhost ip6-loopback
    fe00::0 ip6-localnet
    ff00::0 ip6-mcastprefix
    ff02::1 ip6-allnodes
    ff02::2 ip6-allrouters
    172.17.0.2      561e50dcbe9d

# 宿主机ip一般为：172.17.0.1
# 宿主机查看网络：
docker network list
    NETWORK ID     NAME                       DRIVER    SCOPE
    edbefcc2d9c0   bigchaindb_default         bridge    local
    861f9024a52c   bridge                     bridge    local    可以通过 docker container inpect 查看到当前容器在默认的bridge里
    576ba26dc8a2   datatransferapi_default    bridge    local
    d7fbe1ff94a1   docker-hive_default        bridge    local
    dbf81e4760a7   host                       host      local
    7c939b61da9d   none                       null      local
    e228de87cd73   sentry_onpremise_default   bridge    local

docker network inspect bridge
    [
        {
            "Name": "bridge",
            "Id": "861f9024a52c4871deac1aa778382481cd26f22872da8a708df46454c7a3b186",
            "Created": "2022-09-23T01:10:50.769410103Z",
            "Scope": "local",
            "Driver": "bridge",
            "EnableIPv6": false,
            "IPAM": {
                "Driver": "default",
                "Options": null,
                "Config": [
                    {
                        "Subnet": "172.17.0.0/16",
                        "Gateway": "172.17.0.1"    # 宿主机网关的地址
                    }
                ]
            },
            "Internal": false,
            "Attachable": false,
            "Ingress": false,
            "ConfigFrom": {
                "Network": ""
            },
            "ConfigOnly": false,
            "Containers": {
                "561e50dcbe9d58e2935883feaee91c741f9c26a9424c8b43736d37d9b3d9ecbc": {
                    "Name": "funny_fermat",
                    "EndpointID": "3782dd31dcb0c97b5d848482fbf8b9897728a2b144d805fe9eb5c1027a98be0b",
                    "MacAddress": "02:42:ac:11:00:02",
                    "IPv4Address": "172.17.0.2/16",   # 这里便是容器的网址
                    "IPv6Address": ""
                }
            },
            "Options": {
                "com.docker.network.bridge.default_bridge": "true",
                "com.docker.network.bridge.enable_icc": "true",
                "com.docker.network.bridge.enable_ip_masquerade": "true",
                "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
                "com.docker.network.bridge.name": "docker0",
                "com.docker.network.driver.mtu": "1500"
            },
            "Labels": {}
        }
    ]

# 想访问不在同一网络下的但是映射了端口在宿主机13306端口的mysql-wsl容器便可以使用： 172.17.0.1:13306
```

### 记录


go标准库不在GOPATH下，而是在GOROOT下。
具体的位置：/usr/local/go/src

GOPAT文件夹有点类似虚拟环境的作用，把同一项目的依赖都放在此目录下。
公用的标准库，仍然放在GOROOT下。
这与go是编译型语言有关系吧，只要控制好编译时，包路径就好了。

go env


