package combination_test

import (
	"design/structure/combination"
	"fmt"
	"testing"
)

// 现在
// d1 下面有子部门d2并且有两个直属员工e1,e2
// d2 下面有子部门d3并且有一个直属员工e3
// d3 下面有直属员工e4,e5
func TestCombination(t *testing.T) {
	e1 := combination.CreateEmployee(1500)
	e2 := combination.CreateEmployee(2000)
	e3 := combination.CreateEmployee(3000)
	e4 := combination.CreateEmployee(2000)
	e5 := combination.CreateEmployee(3000)

	d1 := combination.CreateDepartment()
	d2 := combination.CreateDepartment()
	d3 := combination.CreateDepartment()

	d1.AddSubNode(d2)
	d1.AddSubNode(e1)
	d1.AddSubNode(e2)

	d2.AddSubNode(d3)
	d2.AddSubNode(e3)

	d3.AddSubNode(e4)
	d3.AddSubNode(e5)

	fmt.Println("d1 部门总薪水:", d1.CalcSalary())
	fmt.Println("d1 部门及其子部门的人数总和", d1.CountNumber())

}
