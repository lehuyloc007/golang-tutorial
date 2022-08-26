package main

import "fmt"

func main17() {
	/*
		--- Defer: là một buid-in function của go ---
		Defer dùng để trì hoãn thực thi một hàm đến khi các hàm xung quanh nó được return
		Khi gọi nhiều Defer thì nó sẽ đẩy các hàm có gọi defer vào 1 list
	*/

	/*
		--- Panic: là một buid-in function của go ---
		Panic dùng để dừng flow đang chạy của đoạn code và nó bắt đầu quá trình panicking
		Khi một hàm F call Panic thì những cái thực thi của F bị dừng lại và những cái defer function của F được thực thi
	*/

	/*
		--- Recover: là một buid-in function của go ---
		Recover được gọi dùng để láy lại quyền điều khiển của panicking. Nó sẽ hữu dụng khi dùng trong 1 hàm có defer function
	*/

	f()
	fmt.Println("Returned normally from f.")

}
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i+1))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
