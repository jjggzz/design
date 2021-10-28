package factory

import "fmt"

// 定义一个父接口Person,所有的子类都实现了此接口,通过工厂可创建子该接口的所有实现
type Person interface {
	Speck()
}

/**
 定义了小学生\中学生\大学生三个结构体
 都实现了Person接口
**/
type primaryStudent struct{}

func (primaryStudent) Speck() {
	fmt.Println("PrimaryStudent speck")
}

type middleStudent struct{}

func (middleStudent) Speck() {
	fmt.Println("MiddleStudent speck")
}

type collegeStudent struct{}

func (collegeStudent) Speck() {
	fmt.Println("CollegeStudent speck")
}

/**
 定义了小学老师\中学老师\大学老师三个结构体
 都实现了Person接口
**/
type primaryTeacher struct{}

func (primaryTeacher) Speck() {
	fmt.Println("primaryTeacher speck")
}

type middleTeacher struct{}

func (middleTeacher) Speck() {
	fmt.Println("middleTeacher speck")
}

type collegeTeacher struct{}

func (collegeTeacher) Speck() {
	fmt.Println("collegeTeacher speck")
}
