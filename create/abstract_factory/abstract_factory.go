package abstract_factory

// 定义了创建汽车配件的工厂接口
type CarPartsFactory interface {
	MakeEngine() Engine
	MakeWheel() Wheel
}

// 本田产商实现了生产配件的工厂，分别生产本田的引擎和本田的轮子
type HondaCarPartsFactory struct{}

func (HondaCarPartsFactory) MakeEngine() Engine {
	return HondaEngine{}
}

func (HondaCarPartsFactory) MakeWheel() Wheel {
	return HondaWheel{}
}

// 奔驰产商实现了生产配件的工厂，分别生产奔驰的引擎和奔驰的轮子
type BenzCarPartsFactory struct{}

func (BenzCarPartsFactory) MakeEngine() Engine {
	return BenzEngine{}
}

func (BenzCarPartsFactory) MakeWheel() Wheel {
	return BenzWheel{}
}
