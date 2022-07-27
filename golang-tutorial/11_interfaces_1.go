package main

import "fmt"

/*---Ví dụ dử dựng interface để tính tiền lương của từng nhân viên---*/

//khai báo interface
type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

//tiền lương của nhân viên permanent bằng tổng của basic pay và pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//tiền lương của nhân viên contract chỉ là basic pay
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
tổng chi phí được tính bằng cách duyệt qua từng phần tử của slice SalaryCalculator
và tính tổng mức lương của từng nhân viên
*/
//Khai báo hương thức của interface
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}
func main11_1() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	//Ưu điểm lớn nhất của hàm totalExpense này là nó có thể được mở rộng đến bất kỳ loại nhân viên mới nào mà không cần phải thay đổi code.
	//Giả sử công ty bổ sung một loại nhân viên mới là Freelancer với cách tính lương khác.
	//Freelancer này chỉ việc truyền vào đối số slice của hàm totalExpense mà không phải thay đổi bất kỳ 1 dòng code nào trong hàm totalExpense.
	//Freelancer cũng implement interface SalaryCalculator.
	totalExpense(employees)
}
