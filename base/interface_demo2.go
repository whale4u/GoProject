package main

import (
	"fmt"
)

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

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// 这样做最大的优点是：totalExpense 可以扩展新的员工类型，而不需要修改任何代码。
// 假如公司增加了一种新的员工类型 Freelancer，它有着不同的薪资结构。
// Freelancer只需传递到 totalExpense 的切片参数中，无需 totalExpense 方法本身进行修改。
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 3000, 20}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
