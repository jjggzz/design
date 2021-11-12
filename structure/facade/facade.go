package facade

import "fmt"

// 有一个画图的抽象
// 并且有两个实现类
type Shape interface {
	Draw()
}

type Circle struct{}

func (Circle) Draw() {
	fmt.Println("Circle draw")
}

type Rectangle struct{}

func (Rectangle) Draw() {
	fmt.Println("Rectangle draw")
}

// 一个门面,持有了两个具体的实现,并暴露了两个方法,
// 在方法中直接调用所持有实现的具体方法
type ShapeMake struct {
	C Shape
	R Shape
}

// 暴露一个可以简单创建对象的方法
func CreateShapeMake() ShapeMake {
	return ShapeMake{
		C: Circle{},
		R: Rectangle{},
	}
}

func (sm ShapeMake) ShapeCircle() {
	sm.C.Draw()
}

func (sm ShapeMake) ShapeRectangle() {
	sm.R.Draw()
}
