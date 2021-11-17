package strategy

import "fmt"

// 定义了一个旅行的接口
type Travel interface {
	run()
}

// 可以选择坐小汽车旅行
type Car struct{}

func (Car) run() {
	fmt.Println("坐小汽车旅行")
}

// 可以选择坐火车旅行
type Train struct{}

func (Train) run() {
	fmt.Println("坐火车旅行")
}

// 可以选择飞机汽车旅行
type Plane struct{}

func (Plane) run() {
	fmt.Println("坐飞机旅行")
}

// 人只需持有一个具有旅行行为的实例即可
// 通过什么方式去旅行它并不关心
type Person struct {
	T Travel
}

func (p Person) ToRun() {
	p.T.run()
}
