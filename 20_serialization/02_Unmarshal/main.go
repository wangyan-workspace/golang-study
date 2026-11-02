package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func main() {
	//json字符串
	var str = `{"ID":34,"Gender":"男","Name":"Tom","Sno":"20260001"}`
	var s Student
	err := json.Unmarshal([]byte(str), &s)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", s) // main.Student{ID:34, Gender:"男", Name:"Tom", Sno:"20260001"}
	fmt.Println(s.Name)    // Tom
}
