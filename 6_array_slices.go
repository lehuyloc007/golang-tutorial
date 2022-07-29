package main

import (
	"fmt"
)

type Vertex1 struct {
	Lat, Long float64
}

var q map[string]Vertex1

//cách khai báo2(map) khai báo và cấp phát luôn
var z = map[string]Vertex1{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.68433, -122.39967,
	},
}

//cách khai báo3(map) khai báo và cấp phát luôn nhưng ko cần Vertex trong item
//var z = map[string]Vertex1{
//"Bell Labs": {
//40.68433, -74.39967,
//},
//"Google": {
//37.68433, -122.39967,
//},
//}

func main6() {
	//Array trong Go thì là mảng cố định không hơn ko kém nên không thể thay đổi thêm bớt thành phần trong đó
	//Nguyên tắc của mảng length là cố định
	//Bản chất khi ta thêm hoặc bớt phần tử vào mảng là nó copy mảng cũ ra mảng mới với số lượng phần tử lớn hơn mảng cũ và xoá hoặc thêm vào mảng mới
	var a [2]string //Khai báo mảng có 2 phần tử dạng string
	//Cách khai báo trước và gán sau
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
	//Range khai báo1: for biến1, biến2 := range mảng cần lặp
	// trong đó biến 1 là index biến 2 là value
	for o, p := range n {
		fmt.Printf("2**%d = %d\n", o, p)
	}

	//Range khai báo2: for biến1 := range mảng cần lặp
	// trong đó biến 1 là index
	//for o:= range n {
	//fmt.Printf("%d\n", o)
	//}

	//Range khai báo3: for _, biến1 := range mảng cần lặp
	// trong đó biến 1 là value
	//for _, o:= range n {
	//fmt.Printf("%d\n", o)
	//}

	//MAP là key value
	//cấp phát bằng make
	//cách khai báo1(map)
	q = make(map[string]Vertex1)
	q["Bell Labs"] = Vertex1{
		40.68433, -74.39967,
	}
	fmt.Println(q["Bell Labs"])

	fmt.Println(z)
	//cập nhật hoặc thêm mới element trong map
	z["Bell Labs"] = Vertex1{10, 30}
	fmt.Println(z)
	//xóa 1 element trong map
	delete(z, "Bell Labs")
	fmt.Println(z)
	//kiểm tra key có xuất hiện hay không bằng phép gán hai giá trị
	//khi mà tra kiểm tra key mà ko có trong map thì value trả về bằng 0 mà ta ko biết nó va lue nó có là 0 hay 0 nên tra phải dùng cách này để kiểm tra
	elem, ok := z["Bell Labs"]
	fmt.Println("The value:", elem, "Present?", ok)
}
