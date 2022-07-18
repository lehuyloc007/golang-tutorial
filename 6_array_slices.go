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

	//Để giải quyết vấn đề nguyên tắc của mảng là length cố định thì ta có slices
	//Slices tương tự như mảng, nhưng mạnh mẽ và linh hoạt hơn.
	//Giống như mảng, slices cũng được sử dụng để lưu trữ nhiều giá trị cùng kiểu trong một biến duy nhất.
	//Tuy nhiên chiều dài của slices có thể lớn hơn và nhỏ hơn

	//Kb1: Khai báo 1 slices rỗng
	var myslice []int // ?? khai báo myslice := []int lỗi
	//slices là con trỏ vậy nên giá trị mặc định của nó là nil
	var h []int
	if h == nil {
		fmt.Println("nil!")
	}
	//i là con trỏ nên cũng có kiểu nil
	var i *int
	if i == nil {
		fmt.Println("nil!")
	}

	fmt.Println(myslice)
	//Kb2: Khai báo 1 sliecs và gán giá trị cho nó
	c := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(c)

	//Slices là con trỏ trỏ tới tới mảng
	// khai báo và lấy đến phần tử thứ 4 kể từ phần tử đầu tiên
	d := c[:4] //cách 2 var d []int = primes[:4]
	fmt.Println(d)

	// khai báo và lấy ra phần tử ở vị trí thứ 1 đến phần tử thứ 4 thì dừng lại
	e := c[1:4] //cách 2 var e []int = primes[1:4]
	fmt.Println(e)

	// khai báo và lấy ra phần tử từ vị trí thứ 3
	f := c[3:] //cách 2 e []int = primes[3:]
	fmt.Println(f)

	//SLICES bản chất là cấu trúc dữ liêu. Có thể gọi kiểu này là mảng
	//Bên trong nó là 1 cái mảng , length(len) và capacity(cap)
	// Trong đó: len() là độ chứa hiện tại, cap() là độ chứa tối đa
	g := []int{2, 3, 4, 5, 9, 7, 8}
	fmt.Printf("len=%d cap=%d %v\n", len(g), cap(g), g)
	g = g[:4]
	fmt.Printf("len=%d cap=%d %v\n", len(g), cap(g), g)

	//Kb3: Khai báo slice cùng với make name := make([]datatype, len, cap)
	//cách1: khai báo slice có len và cap bằng nhau
	k := make([]int, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(k), cap(k), k)
	//cách1: khai báo slice có len và cap khác nhau
	l := make([]int, 2, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(l), cap(l), l)

	//slice of slice: là mảng 2 chiều

	//thêm phần tử vào slices
	//append(mảng cần thêm, ...các phần tử cần thêm vào mảng)
	m := []int{}
	fmt.Printf("len=%d cap=%d %v\n", len(m), cap(m), m)
	m = append(m, 1, 2, 3, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(m), cap(m), m)

	//RANGE: thường đi với for dùng để lấy ra từng phần của mảng
	n := []int{1, 2, 4, 8, 16, 32, 64, 128}
	//Range khai báo1: for biến1, biến2 := range mảng
	for o, p := range n {
		fmt.Printf("2**%d = %d\n", o, p)
	}
}
