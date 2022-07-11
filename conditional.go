package main

import (
	"fmt"
	"runtime"
)

func main3() {

	//For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//trong Go không có white
	//If
	x := 2
	if x >= 2 {
		fmt.Println(x)
	}
	//If câu lệnh phối hợp vừa gán vừa kết hợp với điều kiện
	if v := 3; v > 2 {
		fmt.Println(v)
	} else {
		fmt.Println(v - 1)
	}
	//Switch cũng có câu lệnh gán kết hợp với điều kiện
	switch os := runtime.GOOS; os {
	case "darwwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("%S. \n", os)
	}
	// Defer là 1 từ khóa chỉ có Go mới có.
	// Là từ khóa mà thực chỉ thi phần phía sau của của nó sau khi 1 funcion kết thúc
	// Defer mang 1 đăc tính quan trọng là stacking(xếp trồng lên nhau) last in first out
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
