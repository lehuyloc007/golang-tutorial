package main

import (
	"fmt"
)

// function cũng coi là 1 giá trị truyền vào hàm
// còn gọi là con trỏ hàm
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func main7() {

	fmt.Println("hello")
}
