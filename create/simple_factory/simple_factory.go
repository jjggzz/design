package simple_factory

import "fmt"

// 定义一个接口
type Person interface {
	Speck()
}

// 定义一个student结构体实现该接口的所有方法,这个实现对外部不可见
type student struct {
}

func (student) Speck() {
	fmt.Println("Student speck")
}

// 定义一个teacher结构体实现该接口的所有方法,这个实现对外部不可见
type teacher struct {
}

func (teacher) Speck() {
	fmt.Println("Teacher speck")
}

// 提供一个对外暴露的方法或者函数，通过传入的参数创建对应的实现
func CreatePerson(typeName string) Person {
	switch typeName {
	case "student":
		return student{}
	case "teacher":
		return teacher{}
	default:
		return nil
	}
}
