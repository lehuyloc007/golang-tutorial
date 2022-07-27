package main

import (
	"fmt"
	"math"
)

// function cũng coi là 1 giá trị truyền vào hàm
// còn gọi là con trỏ hàm
// hàm compute sẽ nhận vào 1 function và thực thi function đó
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func main7() {
	//khỏi tạo hypot là 1 function
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	//in ra kết quả của hypot
	fmt.Println(hypot(5, 12))

	//in ra kết quả khi thực hiện
	fmt.Println(compute(hypot))

	//math.Pow cũng nhận 2 tham số là float64 giống hàm fn nên compute nó hiểu được
	fmt.Println(compute(math.Pow))
}
