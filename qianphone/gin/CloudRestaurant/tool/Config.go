package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

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

// 包内私有变量
var _config *Config

// 返回解析好的配置数据
func GetAppConfig() *Config {
	return _config
}

// 解析配置文件
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
