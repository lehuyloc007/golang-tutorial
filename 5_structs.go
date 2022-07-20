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
	//type [tên_struct] struct {
	// trường của struct (fields)
	//}
	//Cách 1:
	v := MyStruct{1, 2}
	//thay đổi giá trị X của structs
	v.X = 4

	//Cách 2
	n := MyStruct{X: 1, Y: 2}
	//cách 3
	emloyee := struct {
		names string
		age   int
	}{"hello", 30}

	//struct trong struct
	//không phải là chỉ một cách nhúng trường struct trong struct, chúng ta cùng có thể nhúng những type array như kiểu dữ liệu string, int, bool...
	type Skills []string
	type Person struct {
		name string
		age  int
		sex  bool
	}
	type Employee struct {
		Person
		Skills
		salary int
		name   string
	}
	employee := Employee{Person: Person{name: "A", age: 30, sex: true}, Skills: Skills{"code", "tester"}, salary: 600, name: "Samnang"}
	fmt.Println(employee.name)
	fmt.Println(employee.Person.name)

	fmt.Println(v.X, v.Y)
	fmt.Println(n)
	fmt.Println(emloyee)
}
