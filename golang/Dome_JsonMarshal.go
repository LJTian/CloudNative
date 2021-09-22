package main

import (
	"encoding/json"
	"fmt"
)

// Json 编解码 (解析各个字段)
type JsonMarshal struct {
	Msg  string // 消息
	Code string // 状态码
}

func UnMarshalAny(str []byte) {
	var obj interface{}

	if err := json.Unmarshal(str, &obj); err != nil {
		fmt.Println("Unmarshal err")
	}
	objMap, _ := obj.(map[string]interface{})
	for k, v := range objMap {
		switch value := v.(type) {
		case string:
			fmt.Printf("type of %s is string, value is %v \n", k, value)
		case interface{}:
			fmt.Printf("type of %s is interface{}, value is %v \n", k, value)
		default:
			fmt.Printf("type of %s is wrong, value is %v \n", k, value)
		}
	}
}

func MarshalMain() {

	Msg := JsonMarshal{"成功", "00"}
	var bytes []byte
	bytes, _ = json.Marshal(&Msg)
	UnMarshalAny(bytes)
}
