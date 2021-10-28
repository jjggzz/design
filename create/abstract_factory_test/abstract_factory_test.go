package abstractfactory_test

import (
	"design/create/abstract_factory"
	"testing"
)

// 由于一台汽车需要引擎和轮子才能跑，所以需要先创建引擎和轮子
// 工厂模式(factory)也能够通过创建两个工厂(引擎工厂和轮子工厂)来实现对象的创建
// 但是工厂模式创建的对象没有关联性，在使用中本田的轮子可能和奔驰的引擎一起构建了汽车对象
// 假如两种不同产商的产品无法一起协同工作的话就会有问题，使用者也需要了解它们之间的关系才能够放心使用

// 而抽象工厂(abstractFactory)将一组抽象的实现聚集在一起。就列子而言。需要创建一台汽车时
// 先创建好一个产商的工厂，通过它创建该产商提供的所有汽车需要的配件，最终实现汽车的创建。
// 而这些配件之间的兼容性由工厂的创建者保证，使用者就不需要了解它们之间是否能够正常的协作

// 抽象工厂的缺点在于，如果新加入一个配件，例如挡风玻璃，
// 需要为所有工厂都维护对应的实现(就算挡风玻璃与任何类型零件都能一起工作)
// 假如不这么做，有些工厂就无法创建小汽车所需要的所有零件了
// 如果告诉使用者某个挡风玻璃是通用的任何零件都能与其一起工作，对于使用者而言又增加了一个复杂度
func TestAbstractFactory(t *testing.T) {
	var hondaFactory abstract_factory.CarPartsFactory = abstract_factory.HondaCarPartsFactory{}

	var car abstract_factory.Car = abstract_factory.Car{
		Engine: hondaFactory.MakeEngine(),
		Wheel:  hondaFactory.MakeWheel(),
	}

	car.Engine.Run()

}
