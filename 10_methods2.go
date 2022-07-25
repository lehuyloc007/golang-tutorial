package main

import (
	"fmt"
)

type AD struct {
	name string
	age  int
}

/*
Phương thức với vật nhận là giá trị
*/
func (e AD) changeName(newName string) {
	e.name = newName
}

/*
Phương thức với vật nhận là con trỏ
*/
func (e *AD) changeAge(newAge int) {
	e.age = newAge
}

func main10_2() {
	//-Pointer receivers: con trỏ của receivers(vật nhận)
	//-Ở các ví dụ trước các phương thức có receivers là giá trị. Tuy nhiên ta có thể tạo các phương thức với receivers là con trỏ.
	//-Điểm khác nhau giữa receivers là giá trị và receivers là con trỏ là:
	//Những thay đổi được thực hiện bên trong một phương thức đối với receivers là con trỏ sẽ thực sự thay đổi receivers đó, còn đối với receivers là giá trị thì không

	e := AD{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	//có thể viết e.changeAge(51) ~ (&e).changeAge(51) vì lúc này go đã thấy receivers là con trỏ nên ta có thể viết luôn như vậy
	(&e).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}
