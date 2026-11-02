package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"` //通过指定 tag 实现 json 序列化该字段时的 key
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {
	var s = Student{
		ID:     34,
		Gender: "男",
		Name:   "Tom",
		Sno:    "20260001",
	}
	fmt.Printf("%#v\n", s) // main.Student{ID:34, Gender:"男", Name:"Tom", Sno:"20260001"}

	jsonByte, _ := json.Marshal(s)
	jsonStr := string(jsonByte)
	fmt.Printf("%v", jsonStr) // {"id":34,"gender":"男","name":"Tom","sno":"20260001"}
}
