package main

import (
	"fmt"
)

func main4() {
	//pointer: con trỏ
	i, j := 42, 2071
	// khi khai báo 1 biến có giá trị thì nó sẽ nằm trên 1 ô nhớ trên ram nó sẽ có địa chỉ
	// &i là cách lấy ra địa chỉ của ô nhớ
	p := &i //point to i (p là poiter trỏ đến điạ chỉ của i)
	//*p là lấy giá trị của pointer đó ra
	fmt.Println(*p)
	*p = 21 //set giá trị của pointer = 21
	fmt.Println(i)
	p = &j       // thay đổi ponter đến địa chỉ của i
	*p = *p / 37 // thay đổi giá trị của pointer
	fmt.Println(j)

	// Tại sao lại phải dùng con trỏ vì trong Go có 2 loại tham chiếu và tham trị
	// Khi gán i = 42 thì ta gọi đó là tham trị
	// Nếu lấy k = i thì nó copy value sau đó gán k = 10 thì i ko đổi gọi đó là tham chiếu
	// Vậy Con trỏ dùng để Thao tác với địa chỉ bằng các phép tính toán với số được lưu trong nó và Thao tác với giá trị tại địa chỉ mà nó lưu mà thôi.
	// Như vậy lợi ích con trỏ: Tiết kiệm memory
}
