

### 功能点实现记录

1. 多模板

[Gin 默认允许只使用一个 html 模板。 查看多模板渲染 以使用 go 1.6 block template 等功能。](https://gin-gonic.com/zh-cn/docs/examples/multiple-template/)
[example.go](https://github.com/gin-contrib/multitemplate/blob/master/example/advanced/example.go)


2. 结构体初始化问题
```go
    // var config *Config
    // decoder.Decode(_config) 会报错 panic: json: Unmarshal(nil *tool.Config)  

    // 1-结构体没有默认值，需要使用new初始化，
    // 2-初始化后的结构体才能被修改
    config := new(Config)
    if err := decoder.Decode(config); err != nil {
        panic(err)
    }
```

4. redis链接数查看
```sh
redis-wsl:0>info clients
    "# Clients
    connected_clients:6
    cluster_connections:0
    maxclients:10000
    client_recent_max_input_buffer:56
    client_recent_max_output_buffer:0
    blocked_clients:0
    tracking_clients:0
    clients_in_timeout_table:0
```


6. gin json格式响应数据
```go
// https://wnanbei.github.io/post/gin-%E8%BF%94%E5%9B%9E%E5%93%8D%E5%BA%94%E6%96%B9%E5%BC%8F/

// 返回普通 JSON 数据
func (c *Context) JSON(code int, obj interface{})
// 基于 JSON，解决浏览器跨域访问问题
func (c *Context) JSONP(code int, obj interface{})
// 按字面对字符进行编码，不使用 unicode 替换特殊 HTML 字符转义
func (c *Context) PureJSON(code int, obj interface{})
// 把非 Ascii 字符串转为 unicode 编码
func (c *Context) AsciiJSON(code int, obj interface{})
// 防止 JSON 劫持
func (c *Context) SecureJSON(code int, obj interface{})
// 返回缩进美化后的 JSON 数据
func (c *Context) IndentedJSON(code int, obj interface{})

```