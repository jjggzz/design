package combination

// 现在有部门和员工,部门下面可能有子部门也有直属员工
// 怎样用统一的一种方式来处理 部门和员工两种不同的对象呢?
// 比如说计算工资,计算人数等等

// 抽象一个节点,这个节点代表的可能是员工也可能是部门
type Node interface {
	CalcSalary() float32
	CountNumber() int
}

// 为了简单使用,提供一个简单的创建器
func CreateEmployee(salary float32) *Employee {
	return &Employee{
		salary: salary,
	}
}

type Employee struct {
	salary float32
}

// 对于员工而言,计算薪水就是直接返回它的工资
func (e *Employee) CalcSalary() float32 {
	return e.salary
}

// 对于员工而言,计算人数就是直接返回1
func (e *Employee) CountNumber() int {
	return 1
}

// 为了简单使用,提供一个简单的创建器
func CreateDepartment() *Department {
	return &Department{
		subNodeList: make([]Node, 0),
	}
}

type Department struct {
	subNodeList []Node
}

// 对于部门而言可以为它添加子节点
func (d *Department) AddSubNode(node Node) {
	d.subNodeList = append(d.subNodeList, node)
}

// 对于部门而言,计算薪水就是返回它子节点的薪水的和
func (d *Department) CalcSalary() float32 {
	var count float32 = 0
	for _, v := range d.subNodeList {
		count += v.CalcSalary()
	}
	return count
}

// 对于部门而言计算人数就是直接返回所有子节点的人数的和
func (d *Department) CountNumber() int {
	var count int = 0
	for _, v := range d.subNodeList {
		count += v.CountNumber()
	}
	return count
}
