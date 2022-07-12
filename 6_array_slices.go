package main

import (
	"fmt"
)

func main6() {
	//Array trong Go thì là mảng cố định không hơn ko kém nên không thể thay đổi thêm bớt thành phần trong đó
	//Nguyên tắc của mảng length là cố định
	//Bản chất khi ta thêm hoặc bớt phần tử vào mảng là nó copy mảng cũ ra mảng mới với số lượng phần tử lớn hơn mảng cũ và xoá hoặc thêm vào mảng mới
	var a [2]string //Khai báo mảng có 2 phần tử dạng string
	//Cách khai báo và gán sau
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])

	//Cách khai báo và gán luôn
	b := [3]int{2, 3, 4}
	fmt.Println(b)

	//Slices là con trỏ tới mảng
	//Để giải quyết vấn đề nguyên tắc của mảng là length cố định thì ta có slices

}
