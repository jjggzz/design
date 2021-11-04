package bridge

import "fmt"

// 定义一个画图API接口,具体实现是不同的画图功能
type DrawAPI interface {
	Draw(radius int, x int, y int)
}

type RedPen struct{}

func (RedPen) Draw(radius int, x int, y int) {
	fmt.Println("用红色铅笔画图,radius:", radius, "x:", x, "y:", y)
}

type GreenPen struct{}

func (GreenPen) Draw(radius int, x int, y int) {
	fmt.Println("用绿色铅笔画图,radius:", radius, "x:", x, "y:", y)
}

// 定义一个图形接口,他需要画图
type Shape interface {
	Draw()
}

// 图形的实现类,持有一个画图api接口,这个是可替换的
type Circle struct {
	D      DrawAPI
	Radius int
}

func (c Circle) Draw() {
	c.D.Draw(c.Radius, 0, 0)
}

type Rectangle struct {
	D DrawAPI
	X int
	Y int
}

func (r Rectangle) Draw() {
	r.D.Draw(0, r.X, r.Y)
}
