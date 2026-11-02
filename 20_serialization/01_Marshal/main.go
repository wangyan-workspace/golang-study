package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	name   string //私有属性不能被json包访问
	Sno    string
}

func main() {
	var s = Student{
		ID:     34,
		Gender: "男",
		name:   "Tom",
		Sno:    "20260001",
	}
	fmt.Printf("%#v\n", s) // main.Student{ID:34, Gender:"男", name:"Tom", Sno:"20260001"}

	josonByte, _ := json.Marshal(s)
	/*
		[123 34 73 68 34 58 51 52 44 34 71 101 110 100 101 114 34 58 34 231 148 18334 44 34 83 110 111 34 58 34 50 48 50 54 48 48 48 49 34 125]
	*/
	fmt.Println(josonByte)
	jsonStr := string(josonByte)
	fmt.Printf("%v", jsonStr) // {"ID":34,"Gender":"男","Sno":"20260001"}

}
