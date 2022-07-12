package main

import (
	"fmt"
)

type MyStruct struct {
	X int
	Y int
}

func main5() {
	//structs: là cấu trúc dữ liệu
	// structs: là kiểu dữ liệu tự định nghĩa được cấu tạo bỏi các flields

	//khỏi tạo structs
	//Cách 1:
	v := MyStruct{1, 2}
	//thay đổi giá trị X của structs
	v.X = 4

	//Cách 2
	n := MyStruct{X: 1, Y: 2}

	fmt.Println(v.X, v.Y)
	fmt.Println(n)
}
