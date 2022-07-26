package main

import (
	"fmt"
)

/*---Buffered Channel---*/
//Buffered Channel là một channel trong Golang có khả năng lưu trữ được dữ liệu bên trong nó.
//Khả năng này được mô tả như sức chứa (capacity) của channel.

//Khai báo: chanName := make(chan Type, capacity)
//Trong đó: Type là kiểu dữ liệu, capacity sức chứa

func main15() {
	/*---Các đặc tính của Buffered Channel---*/
	//Buffered Channel có len và cap
	//len: là số lượng giá trị/dữ liệu hiện đang có trong buffered channel.
	//cap: là sức chứa tối đa của buffered channel.
	bufferedChan := make(chan int, 5)

	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))

	bufferedChan <- 1
	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))

	bufferedChan <- 2
	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))

	<-bufferedChan
	fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))

	// Buffered Channel sẽ block goroutine hiện tại nếu vượt sức chứa
	// bufferedChan1 := make(chan int, 5)
	// bufferedChan1 <- 1
	// bufferedChan1 <- 2
	// bufferedChan1 <- 3
	// bufferedChan1 <- 4
	// bufferedChan1 <- 5
	// fmt.Printf("BufferChan1 has len = %d, cap = %d\n", len(bufferedChan1), cap(bufferedChan1))
	// bufferedChan1 <- 6 // deadlock here code bị lỗi

	//Lấy dữ liệu từ empty buffered channel sẽ block goroutine
	//Tương tự với việc đẩy dữ liệu vào một full buffered channel, việc lấy dữ liệu từ một empty bufffered channel cũng block goroutine hiện tại.
	//bufferedChan := make(chan int, 5)
	//fmt.Printf("BufferChan has len = %d, cap = %d\n", len(bufferedChan), cap(bufferedChan))
	//<-bufferedChan // deadlock here

	//Lưu trữ dữ liệu theo thứ tự FIFO (First-In-First-Out)
	//Dữ liệu đẩy vào và lấy ra khỏi buffered channel theo thứ tự FIFO, việc này khiến nó hoạt động như một queue (hàng đợi):
	// bufferedChan := make(chan int, 5)

	// for i := 1; i <= 5; i++ {
	// 	bufferedChan <- i
	// }

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(<-bufferedChan)
	// }

	/*---Khác biệt giữa buffered channel và unbuffered channel(channel)---*/
	//Sự khác biệt lớn nhất giữa buffered channel và unbuffered channel đó là về capacity.
	//Buffered channel sẽ yêu cầu khi báo capacity lúc khởi tạo channel, unbuffer channel thì không cần.
	// bufferedChan := make(chan int, 1)
	// unbufferedChan := make(chan int)

	// bufferedChan <- 1   // OK
	// unbufferedChan <- 1 // deadlock
}
