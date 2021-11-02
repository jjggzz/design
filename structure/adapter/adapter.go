package adapter

import "fmt"

// 销售员具有计算价格的能力
type Salesclerk interface {
	CalcPrice()
}

type SalesclerkImpl struct{}

func (SalesclerkImpl) CalcPrice() {
	fmt.Println("calc price")
}

// 店长除了有计算价格的能力外还能进行管理
// 计算价格的操作都是一样的
type ShopManager interface {
	Manager()
	CalcPrice()
}

// 现在有一个店长的适配器,它实现了店长的所有操作
// 内部持有一个销售员的对象,这样就可以使用销售员计算价格的方法了
// 但是管理的方法还需要自己实现
type ShopManagerAdapter struct {
	Salesclerk Salesclerk
}

func (s ShopManagerAdapter) Manager() {
	fmt.Println("do manager")
}

func (s ShopManagerAdapter) CalcPrice() {
	s.Salesclerk.CalcPrice()
}
