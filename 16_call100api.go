package main

import (
	"fmt"
	"time"
)

func main16() {
	fmt.Println("Start")
	//giả sử có 1000 url cần gọi thì ta cần tạo số lượng request max nhất là 1000
	numberUrlRequest := 1000

	//tạo ra channel có thể gửi vào 1000 dữ liệu
	quereChan := make(chan int, numberUrlRequest)

	//khai báo số lượng worker
	maxWorkerNumber := 5
	//(2)Mở rộng thêm: cách tạo thêm 1 channel để thấy được 5 cái worker đã dùng khi đó ta sẽ không phải dùng sleep 2s
	//(2)khai báo ra channel sau đó sẽ được nhận vào là các worker
	doneChan := make(chan string)

	//tạo ra 5 worker là 5 goroutine
	for i := 0; i < maxWorkerNumber; i++ {
		go func(nameWorker string) {
			//chúng ta không thể truyền thẳng i vào nameWorker trong hàm crawlUrl vì biến i đã được sử dụng ở hàm for range đây là mặc định
			// nên ta phải khai báo 1 tham số nameWorker ở hàm arrow func và truyền i vào đó khi nó chạy
			for valueQuereChan := range quereChan {
				//gọi hàm để crawl url
				crawlUrl(nameWorker, valueQuereChan)
			}
			fmt.Printf("%s: is done \n", nameWorker)
			//(2) gửi vào channel doneChan là các worker đang chạy
			doneChan <- nameWorker
		}(fmt.Sprintf("%d", i))
	}

	// gửi dữ liệu vào channel
	for i := 0; i < numberUrlRequest; i++ {
		quereChan <- i
	}

	// dừng lại channel để vòng for range không bị leek khi đó sẽ thấy dòng này chạy fmt.Printf("%s is done: \n", nameWorker)
	close(quereChan)

	for i := 0; i < maxWorkerNumber; i++ {
		//(2) các worker đang chạy ra khỏi channel doneChan
		<-doneChan
	}

	//cho chương trình dừng lại khoảng 2s trước khi thoát
	//time.Sleep(time.Second * 2)

}

//hàm lấy dữ liệu crawl từ url nhận 2 tham số là tên worker và giá trị url cần gọi
func crawlUrl(nameWorker string, valueUrlRequest int) {
	//do crawlUrl được gọi ở 1 vòng for range mà vòng for này chạy liên tục nên nó sẽ ngậm proccess rất lâu nên ta phải cho nó
	//sleep lại 1 chút để cho for này nghỉ lại 1 chút để cho worker swich sang cái khác
	time.Sleep(time.Second / 100)
	fmt.Printf("Name Worker: %s Value Url Request: %d\n ", nameWorker, valueUrlRequest)
}
