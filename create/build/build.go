package build

import "errors"

// 待构建的结构体或对象,这个对象中只有三个属性
// 可能某些场景下属性有十几个之多
type Person struct {
	Name   string
	Age    int
	Gender string
}

// 统一的构建抽象,可以有也可以没有,这里是为了统一抽象
type PersonBuilder interface {
	SetName(name string) PersonBuilder
	SetAge(age int) PersonBuilder
	SetGender(gender string) PersonBuilder
	Build() (*Person, error)
}

// 年轻人的构建器,包含了Person的所有属性
type YoungBuilder struct {
	name   string
	age    int
	gender string
}

func (m *YoungBuilder) SetName(name string) PersonBuilder {
	m.name = name
	return m
}
func (m *YoungBuilder) SetAge(age int) PersonBuilder {
	m.age = age
	return m
}
func (m *YoungBuilder) SetGender(gender string) PersonBuilder {
	m.gender = gender
	return m
}

// 在构建时可以对其中的属性做统一处理或者判断其合法性
// 比如这里要求age必须在0~18之间
// 后面的ElderlyBuilder要求age必须在80~200之间
// 这样就将构建与定义分离开来,可以单独的定义校验逻辑
func (m *YoungBuilder) Build() (*Person, error) {
	if m.name == "" {
		return nil, errors.New("name has not empty")
	}
	if m.age < 0 || 18 < m.age {
		return nil, errors.New("age must between 0 and 18")
	}
	if m.gender != "man" && m.gender != "woman" {
		return nil, errors.New("gender must man or woman")
	}

	return &Person{
		Name:   m.name,
		Age:    m.age,
		Gender: m.gender,
	}, nil
}

type ElderlyBuilder struct {
	name   string
	age    int
	gender string
}

func (e *ElderlyBuilder) SetName(name string) PersonBuilder {
	e.name = name
	return e
}
func (e *ElderlyBuilder) SetAge(age int) PersonBuilder {
	e.age = age
	return e
}
func (e *ElderlyBuilder) SetGender(gender string) PersonBuilder {
	e.gender = gender
	return e
}

func (e *ElderlyBuilder) Build() (*Person, error) {
	if e.name == "" {
		return nil, errors.New("name has not empty")
	}
	if e.age < 80 || 200 < e.age {
		return nil, errors.New("age must between 80 and 200")
	}
	if e.gender != "man" && e.gender != "woman" {
		return nil, errors.New("gender must man or woman")
	}

	return &Person{
		Name:   e.name,
		Age:    e.age,
		Gender: e.gender,
	}, nil
}
