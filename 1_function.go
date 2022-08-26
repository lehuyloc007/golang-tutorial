// package là tên thư mục chứa 1 hoặc nhiều file go trong đó
// tất cả các file chung 1 package thì cái name đều giống nhau
package main

// thư viện dùng để in ra

//khi khai báo thì phải khai báo kiểu của biến truyền vào và kiểu trả về
// có 2 cách khai báo kiểu biến truyền vào
//cách 1: Khai báo từng kiểu của biến
func add1(x int, y int) int {
	return x + y
}

//cách 1; Khai báo chung kiểu của biến
func add2(x, y int, z int) int {
	return x + y + x
}

//multi result là: trả về nhiều kết quả trong 1 function
func swap(x, y string) (string, string) {
	return x, y
}

//name retun values là: trả về nhiều kết quả trong 1 function nhưng có thể khai báo tên của giá trị trả về lúc khai báo function
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// fmt.Println("Hello, World!")

	// fmt.Println(add1(4, 3))

	// fmt.Println(add2(4, 3, 1))

	//có thể khai báo biến để gán cho function bên trong trả về
	// a, b := swap("world", "hello")
	// fmt.Println(a, b)

	// fmt.Println(split(17))

	//trong 1 project go chỉ có 1 function main
	// main2()
	// main3()
	// main4()
	// main5()
	// main6()
	// main7()
	// main16()
	main17()
}
