package employee // 定义该文件属于employee包

import (
	"fmt"
)

// 声明Employee结构体
type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// 将数据和操作数据的方法绑定
func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
