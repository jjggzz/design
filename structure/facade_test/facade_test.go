package facade_test

import (
	"design/structure/facade"
	"testing"
)

// 将多个抽象封装到一个门面中,可以让客户端方便调用
// 使用被简化了
func TestFacade(t *testing.T) {
	var sm facade.ShapeMake = facade.CreateShapeMake()
	sm.ShapeCircle()
	sm.ShapeRectangle()
}
