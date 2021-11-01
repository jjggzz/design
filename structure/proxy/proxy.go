package proxy

import (
	"fmt"
	"time"
)

// 定义行为
type MakeFood interface {
	MakeChicken()
}

// 对象实现,做具体的事情
type MakeFoodImpl struct{}

func (MakeFoodImpl) MakeChicken() {
	fmt.Println("make chicken")
	time.Sleep(time.Second * 1)
}

// 代理对象内部持有一个实现对象并且实现了行为的方法,看起来和被代理的对象类型一样
// 具体的逻辑由实现对象做,代理对象可以在前后做一些跟业务无关的工作
type MakeFoodImplProxy struct {
	MakeFoodImpl
}

func (m MakeFoodImplProxy) MakeChicken() {
	start := time.Now().UnixNano()
	fmt.Println("MakeChicken 调用开始: ", start)
	m.MakeFoodImpl.MakeChicken()
	end := time.Now().UnixNano()
	fmt.Println("MakeChicken 调用结束: ", end)
}
