package bridge_test

import (
	"design/structure/bridge"
	"testing"
)

// 现在可以组合画图API和图形使用了
// 对于要新增颜色,只需要添加DrawAPI接口的实现类即可
// 对于要新增图形,只需要实现Shape接口,并且持有一个DrawAPI
// 即可使用原来已有的所有颜色,原先的代码结构完全不需要变动
func TestBridge(t *testing.T) {
	redDraw := bridge.RedPen{}
	circle1 := bridge.Circle{
		D:      redDraw,
		Radius: 5,
	}
	circle1.Draw()

	greenPen := bridge.GreenPen{}
	circle2 := bridge.Circle{
		D:      greenPen,
		Radius: 5,
	}
	circle2.Draw()

	rectangle := bridge.Rectangle{
		D: greenPen,
		X: 5,
		Y: 6,
	}
	rectangle.Draw()

}
