package main

import (
	"fmt"
)

// function closures
// Một closure là một function value tham chiếu đến các biến từ bên người phần thân của nó
// Hàm có thể truy cập và gán cho các biến được tham chiếu; theo nghĩa này, hàm được "ràng buộc" với các biến.

//function trả về 1 kiểu là function func(int) int nên sẽ phải return giống vậy
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main8() {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}
