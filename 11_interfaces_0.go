package main

import "fmt"

//Trong hướng đối tượng đó là ‘‘Interface xác định hành vi của một đối tượng’’. Nó chỉ xác định những gì đối tượng phải làm
//Interface trong Golang là một kiểu được định nghĩa bởi tập hợp của các method (hàm trong Golang).
//Interface có thể chứa bất kỳ giá trị gì miễn là nó có implement(thực hiện, triển khai) các method này.

//định nghĩa 1 interface
type VowelsFinder interface {
	FindVowels() []rune
}

//Khởi tạo một kiểu dữ liệu tên là MyString.
type MyString string

//MyString implements VowelsFinder
//khởi tạo một interface tên là VowelsFinder có một phương thức là FindVowels() []rune.
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

//Các interface trong Go ngầm hiểu là được implement nếu có một kiểu dữ liệu chứa tất cả các phương thức được khai báo trong interface đó
func main11() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // Điều này được chấp nhận vì MyString implement interface VowelsFinder
	//in ra tất cả các nguyên âm trong chuỗi Sam Anderson
	fmt.Printf("Vowels are %c", v.FindVowels())
}
