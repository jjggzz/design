package abstract_factory

import "fmt"

// 小汽车需要引擎和轮子才能够运行
type Car struct {
	Engine Engine
	Wheel  Wheel
}

// 定义了引擎接口
type Engine interface {
	Run()
}

// 定义了轮子接口
type Wheel interface {
	Go()
}

// 现在有本田的引擎 本田的轮子 奔驰的引擎 奔驰的轮子
type HondaEngine struct{}

func (HondaEngine) Run() {
	fmt.Println("HondaEngine run")
}

type HondaWheel struct{}

func (HondaWheel) Go() {
	fmt.Println("HondaWheel go")
}

type BenzEngine struct{}

func (BenzEngine) Run() {
	fmt.Println("BenzEngine run")
}

type BenzWheel struct{}

func (BenzWheel) Go() {
	fmt.Println("BenzWheel go")
}
