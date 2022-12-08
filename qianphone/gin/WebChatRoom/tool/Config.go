package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName     string      `json:"app_name"`
	AppMode     string      `json:"app_mode"`
	AppHost     string      `json:"app_host"`
	AppPort     string      `json:"app_port"`
	RedisConfig RedisConfig `json:"redis"` // 非匿名字段，注意使用时获取方式
}

// 返回解析好的配置数据
var _config *Config

func GetAppConfig() *Config {
	ParseConfig()
	return _config
}

func ParseConfig() {

	// 读取文件
	file, err := os.Open("./config/app.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	// var config *Config
	// decoder.Decode(_config) 会报错 panic: json: Unmarshal(nil *tool.Config)
	config := new(Config)
	if err := decoder.Decode(config); err != nil {
		panic(err)
	}
	_config = config
}
