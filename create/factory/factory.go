package factory

// 定义了一个Person工厂的接口,该接口的实现类具有创建Person的能力
type PersonFactory interface {
	CreatePerson(string) Person
}

// 定义了Student工厂,他能够根据不同的name创建不同的student
// 它实现了PersonFactory
type StudentFactory struct{}

func (StudentFactory) CreatePerson(name string) Person {
	switch name {
	case "primaryStudent":
		return primaryStudent{}
	case "middleStudent":
		return middleStudent{}
	case "collegeStudent":
		return collegeStudent{}
	default:
		return nil
	}
}

// 定义了Teatcher工厂,他能够根据不同的name创建不同的teatcher
// 它实现了PersonFactory
type TeacherFactory struct{}

func (TeacherFactory) CreatePerson(name string) Person {
	switch name {
	case "primaryTeacher":
		return primaryTeacher{}
	case "middleTeacher":
		return middleTeacher{}
	case "collegeTeacher":
		return collegeTeacher{}
	default:
		return nil
	}
}
