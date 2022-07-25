package main

import "fmt"

/*---Generics---*/
//Generics là một kỹ thuật lập trình thường xuất hiện trong OOP (lập trình hướng đối tượng).
//Kỹ thuật này giúp lập trình viên định nghĩa được những hàm (phương thức hay method/function) chấp nhận các tham số chung chung,
//không cần chỉ định rõ nó thuộc kiểu dữ liệu gì. Tới khi hàm này được sử dụng, người gọi sẽ quyết định việc này.

//Khai báo func nameFunction[T any](s ...T) {
//code
//}
//Trong đó [T any]: kiểu dữ liệu, (s ...T) param truyền vào function
func Print[T any](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
}
func main12() {
	Print("Hello, ", "playground\n")
}
