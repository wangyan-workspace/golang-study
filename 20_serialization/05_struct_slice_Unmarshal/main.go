package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	Id     int
	Gender string
	Name   string
}

// Class 班级
type Class struct {
	Title    string
	Students []Student
}

func main() {
	str := `{"Title":"Go_101","Students":[{"Id":1,"Gender":"男","Name":"stu_1"},{"Id":2,"Gender":"男","Name":"stu_2"},{"Id":3,"Gender":"男","Name":"stu_3"},{"Id":4,"Gender":"男","Name":"stu_4"},{"Id":5,"Gender":"男","Name":"stu_5"},{"Id":6,"Gender":"男","Name":"stu_6"},{"Id":7,"Gender":"男","Name":"stu_7"},{"Id":8,"Gender":"男","Name":"stu_8"},{"Id":9,"Gender":"男","Name":"stu_9"},{"Id":10,"Gender":"男","Name":"stu_10"}]}`

	var c = &Class{}
	err := json.Unmarshal([]byte(str), c)
	if err != nil {
		fmt.Println(err)
	} else {
		/*
		  &main.Class{Title:"Go_101", Students:[]main.Student{main.Student{Id:1, Gender:"男", Name:"stu_1"}, main.Student{Id:2, Gender:"男", Name:"stu_2"}, main.Student{Id:3, Gender:"男", Name:"stu_3"}, main.Student{Id:4, Gender:"男", Name:"stu_4"}, main.Student{Id:5, Gender:"男", Name:"stu_5"}, main.Student{Id:6, Gender:"男", Name:"stu_6"}, main.Student{Id:7, Gender:"男", Name:"stu_7"}, main.Student{Id:8, Gender:"男", Name:"stu_8"}, main.Student{Id:9, Gender:"男", Name:"stu_9"}, main.Student{Id:10, Gender:"男", Name:"stu_10"}}}&main.Class{Title:"Go_101", Students:[]main.Student{main.Student{Id:1, Gender:"男", Name:"stu_1"}, main.Student{Id:2, Gender:"男", Name:"stu_2"}, main.Student{Id:3, Gender:"男", Name:"stu_3"}, main.Student{Id:4, Gender:"男", Name:"stu_4"}, main.Student{Id:5, Gender:"男", Name:"stu_5"}, main.Student{Id:6, Gender:"男", Name:"stu_6"}, main.Student{Id:7, Gender:"男", Name:"stu_7"}, main.Student{Id:8, Gender:"男", Name:"stu_8"}, main.Student{Id:9, Gender:"男", Name:"stu_9"}, main.Student{Id:10, Gender:"男", Name:"stu_10"}}}
		*/
		fmt.Printf("%#v\n", c)

		fmt.Printf("%v", c.Title) // Go_101
	}
}
