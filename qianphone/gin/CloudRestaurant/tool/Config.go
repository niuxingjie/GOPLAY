package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
}

var _config *Config

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
