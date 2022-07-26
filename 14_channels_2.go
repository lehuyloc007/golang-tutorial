package main

import (
	"fmt"
	"time"
)

/*---channels---*/
//Cơ chế block của Channel
//Việc gởi và nhận dữ liệu thông qua Channel sẽ có hỗ trợ cơ chế block.
//Việc này giúp các Goroutines giao tiếp qua Channel một cách đồng bộ (synchronize).
//Về nguyên tắc, Channel sẽ block goroutines nếu nó chưa sẵn sàng.
//Channel giúp các goroutines có thể gởi và nhận được dữ liệu cho nhau một cách an toàn thông qua cơ chế lock-free.
//Channel giúp chúng ta có kênh giao tiếp giữa goroutine an toàn vì: dù cho tốc độ các goroutine khác nhau nhưng khi cần giao tiếp với nhau thì chúng phải dừng lại

func main14_2() {
	myChan := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChan <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 1; i <= 10; i++ {
		fmt.Println(<-myChan)
	}
	//Trong đoạn code trên chúng ta sẽ có 2 goroutine là main và một function vô danh.
	//Main goroutine sẽ lấy dữ liệu từ myChan với tốc độ cao hơn function gởi vào.
	//Ở mỗi vòng lặp, main goroutine cũng sẽ phải đợi cho tới khi myChan có dữ liệu gởi vào.
	//Kết quả ta sẽ thấy các con số được in ra sau mỗi giây.

	//Thật ra đây chính là cơ chế "streaming" rất hữu dụng về sau này. Golang đã làm nó trở nên cực kì gọn và dễ dùng.
}
