package main

import (
	"fmt"
	"math"
)

type AB struct {
	length int
	width  int
}
type AC struct {
	radius float64
}

func (r AB) AB() int {
	return r.length * r.width
}

func (c AC) AB() float64 {
	return math.Pi * c.radius * c.radius
}
func main10_1() {
	//Điểm khác nhau giữa fuction và method:
	//- Go không phải là một ngôn ngữ lập trình hướng đối tượng thuần túy và nó không hỗ trợ các lớp (class).
	//Do đó viết các phương thức trên các kiểu dữ liệu là một cách để có được hành vi tương tự như các lớp.
	//- Các phương thức có thể trùng tên nếu được xác định trên các kiểu dữ liệu khác nhau trong khi các hàm thì không được phép trùng tên.
	r := AB{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.AB())
	c := AC{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.AB())
}
