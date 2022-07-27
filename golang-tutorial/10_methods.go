package main

import (
	"fmt"
	"math"
)

type AA struct {
	X, Y float64
}

func (v AA) FuncAA() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main10() {
	//Phương thức (Method) là một hàm (function) được khai báo cho riêng một kiểu dữ liệu đặc biệt,
	//kiểu dữ liệu này được gọi là vật nhận (receiver) nó được đặt giữa từ khóa func và tên của phương thức.
	//Vật nhận này có thể là kiểu struct (cấu trúc) hoặc non-struct (phi cấu trúc). Vật nhận phải có sẵn để truy cập bên trong phương thức.

	//Cú pháp:
	//func (t Type) methodName(parameter list) {

	//}

	v := AA{3, 4}
	fmt.Println(v.FuncAA())

}
