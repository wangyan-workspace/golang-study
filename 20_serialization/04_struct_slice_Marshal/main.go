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
	c := Class{
		Title:    "Go_101",
		Students: make([]Student, 0),
	}

	// 创建10个学生
	/*
		{Go_101 [{1 男 stu_1} {2 男 stu_2} {3 男 stu_3} {4 男 stu_4} {5 男 stu_5} {6 男 stu_6} {7 男 stu_7} {8 男 stu_8} {9 男 stu_9} {10 男 stu_10}]}
	*/
	for i := 1; i <= 10; i++ {
		s := Student{
			Id:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}

	fmt.Println(c)

	strByte, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	} else {
		strJson := string(strByte)
		/*
			{"Title":"Go_101","Students":[{"Id":1,"Gender":"男","Name":"stu_1"},{"Id":2,"Gender":"男","Name":"stu_2"},{"Id":3,"Gender":"男","Name":"stu_3"},{"Id":4,"Gender":"男","Name":"stu_4"},{"Id":5,"Gender":"男","Name":"stu_5"},{"Id":6,"Gender":"男","Name":"stu_6"},{"Id":7,"Gender":"男","Name":"stu_7"},{"Id":8,"Gender":"男","Name":"stu_8"},{"Id":9,"Gender":"男","Name":"stu_9"},{"Id":10,"Gender":"男","Name":"stu_10"}]}
		*/
		fmt.Println(strJson)
	}
}
