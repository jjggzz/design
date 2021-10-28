package factory_test

import (
	"design/create/factory"
	"testing"
)

/**
	工厂模式(factory)是简单工厂的升级版,
	简单工厂模式(simpleFactory)将所有抽象的实现通过一个方法创建出来
	而工厂模式将抽象的实现分类创建,需要先选好具体的工厂,通过它创建具体的实现
**/
func TestFactory(t *testing.T) {
	// 我们先创建一个能够创建学生的工厂,通过它可以创建不同的学生
	var studentFactory factory.PersonFactory = factory.StudentFactory{}
	var p1 factory.Person = studentFactory.CreatePerson("middleStudent")
	p1.Speck()

	// 创建一个能够创建老师的工厂,通过它可以创建不同的老师
	var teacherFactory factory.PersonFactory = factory.TeacherFactory{}
	var p2 factory.Person = teacherFactory.CreatePerson("middleTeacher")
	p2.Speck()

}
