package main

import "fmt"

/*---Type Assertion---*/
//Type assertion được sử dụng để ép kiểu giá trị cơ bản của interface.
//Cú pháp i.(T) trong đó: i là interface còn T là một kiểu dữ liệu
func main11_3() {
	var i interface{} = "hello"

	//gán s = i có kiểu là string
	s := i.(string)
	fmt.Println(s)

	//cách kiểm tra có thể ép kiểu i về string không ok sẽ trả 2 giá trị true và false sau đó mới gán cho s
	s, ok := i.(string)
	fmt.Println(s, ok)

	//như trường hợp này sẽ báo false nhưng chương trình vẫn chạy tiếp
	f, ok := i.(float64)
	fmt.Println(f, ok)

	//như trường hợp này chương trình sẽ báo lỗi và dừng ngay lập tức
	//f = i.(float64) // panic chương trình bị dừng
	//fmt.Println(f)

	//Type Assertions trong Golang có thể vừa dùng để đưa interface về kiểu dữ liệu bên dưới nó (hoặc một Interface khác nếu thoả điều kiện)
	//vừa dùng để check xem liệu đúng là kiểu dữ liệu muốn ép về hay không.
	//Vì thế, nó có thể trả về một cặp kết quả theo thứ tự: kiểu dữ liệu, giá trị true/false thể hiện cho việc ép kiểu thành công hoặc thất bại.
}
