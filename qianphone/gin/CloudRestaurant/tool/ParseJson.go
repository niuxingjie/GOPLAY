package tool

import (
	"encoding/json"
	"io"
)

// 结构体
type ParseJson struct {
}

// 结构体方法
func (p ParseJson) Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(&v)
}
