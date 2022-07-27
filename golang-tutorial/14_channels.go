package main

import (
	"fmt"
)

/*---channels---*/
//Channel là các kênh giao tiếp trung gian giữa các Goroutines trong Golang.
//Channel giúp các goroutines có thể gởi và nhận được dữ liệu cho nhau một cách an toàn thông qua cơ chế lock-free.

//Khai báo1: var channelName chan Type
//Khai báo2: channelName := make(chan Type)
//Trong đó Type là kiểu dữ liệu

//Gởi dữ liệu vào Channel
//channelName <- value

//Nhận dữ liệu từ Channel
//myVar := <- channelName

func main14() {
	myChan := make(chan int)

	go func() {
		myChan <- 1
	}()

	fmt.Println(<-myChan)
}
