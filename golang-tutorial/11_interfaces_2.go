package main

import "fmt"

/*---Empty Interface---*/
//Một interface không có phương thức nào thì được gọi là interface rỗng
//Vì interface rỗng không có thương thức, nên tất cả các kiểu dữ liệu đều có thể implement interface rỗng.
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func main11_2() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	//Trong ví dụ trên biến i là một interface{}. Vì interface này không hề định nghĩa và yêu cầu bất kỳ một method nào thành ra nó được dùng cho mọi kiểu.
	//Cũng rất dễ hiểu và đảm bảo nguyên tắc implement ngầm định ở trên.

	//Một ứng dụng rất phổ biến với interface{} là sử dụng làm value cho kiểu dữ liệu map của Golang.
	/*
		product := make(map[string]interface{}, 0)

		product["name"] = "Iphone 13 Pro Max"
		product["price"] = 31000000
		product["quantity"] = 40

		fmt.Println(product)
	*/
	//Như vậy ta sẽ có một kiểu map có key luôn là string và value là bất kỳ kiểu dữ liệu nào cũng được.
}
