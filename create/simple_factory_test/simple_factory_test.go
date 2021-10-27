package simplefactory_test

import (
	"design/create/simple_factory"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	// 传入不同的值获取不同的实现
	var student simple_factory.Person = simple_factory.CreatePerson("student")
	student.Speck()

	var teacher simple_factory.Person = simple_factory.CreatePerson("teacher")
	teacher.Speck()
}
