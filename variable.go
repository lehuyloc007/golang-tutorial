package main

import "fmt"

//khai báo nhiều biến cùng 1 kiểu dữ liệu
var c, python, java bool

// khai báo mà muốn gán giá trị luôn thì có 2 cách
//cách 1
var i, j int = 1, 2

// cách 2
var d = "no!"

func main2() {
	// ? Nếu khai báo biến trùng vs biến bên ngoài thì sao
	//zezo value: là khi ta khai báo biến mà không gán giá trị thì nó sẽ nhận giá trị mặc định int 0 float64 0 string ""
	var n int
	fmt.Println(i, n, c, python, java)

	fmt.Println(j, d)

	// cách 3 rút gọn cách này không áp dụng bên ngoài 1 hàm
	e := 3
	fmt.Println(e)

	//ép kiểu nếu ép kiểu từ kiểu bé kên kiểu lớn thì ko sao nhưng nếu từ lớn về nhỏ thì rủi do sai số nếu giá trị lớn hơn kiểu bị ép
	var f int = 42
	var g float64 = float64(f)
	fmt.Println(g)

	//xem kiểu của biến
	fmt.Printf("type of g %T", g)

	//constants là biến được khai báo ra và không thể thay đổi được trong suốt vòng đời của nó và không thể khai báo bằng dấu :=
	const Pi float64 = 3.14

}

/*--- các kiểu dữ liệu của go ---*/
/*
	bool: true false
	string: kiểu chuỗi là dấu nháy đôi
	int int8 int18 int32 int64: kiểu số với kiểu có số đằng sau int32 ~ int thì trước nó là 16 bit âm và 16 bit dương
	uint uint8 uint18 uint32 uint64 unitptr: kiểu số không âm chữ u viết tắt của unsize
	byte là kiểu biến thược kiểu unit8
	rune là kiểu biến int32 đai diện cho một điểm mã unicode
	float32 float64 là kiểu số thực
	complex64 complex128 là kiểu biên phức tạp
*/
