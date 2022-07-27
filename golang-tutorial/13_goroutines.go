package main

import (
	"fmt"
	"time"
)

/*---goroutines---*/
//Goroutines là một trong những tính năng đặc biệt nhất của Golang để lập trình Concurrency(đồng thời) cực kỳ đơn giản.
//https://viblo.asia/p/phan-biet-concurrency-va-parallelism-trong-lap-trinh-gAm5y8PVldb
//Goroutines bản chất là các hàm (function) hay method được thực thi một các độc lập và đồng thời nhưng vẫn có thể kết nối với nhau.
//Một cách ngắn gọn, những thực thi đồng thời được gọi là Goroutines trong Go (Golang).

// Goroutines là lightweight thread được quản lý bởi go runtime

// Cú pháp: go f(x,y,z)

// Goroutines nhẹ hơn thread rất nhiều
// 1 thead có thể quản lý nhiều Goroutines

//trong golang bình thường sẽ chạy async chạy từ trên xuống dưới
//nếu muốn thì mới đặt go vào để chạy Goroutines khác

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main13() {
	go say("world")
	say("hello")
}
