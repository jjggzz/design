package decorate

// 定义公共的饮料抽象,有获取饮料名和计算售价的方法
type Drinks interface {
	GetDrinksName() string
	CalcPrice() float64
}

// 现在有三种基础饮料实现分别实现了饮料的抽象
type Mongo struct{}

func (Mongo) GetDrinksName() string {
	return "芒果茶"
}
func (Mongo) CalcPrice() float64 {
	return 10
}

type Orange struct{}

func (Orange) GetDrinksName() string {
	return "橙汁"
}
func (Orange) CalcPrice() float64 {
	return 7
}

type Coconut struct{}

func (Coconut) GetDrinksName() string {
	return "椰子汁"
}
func (Coconut) CalcPrice() float64 {
	return 15
}

// 现在有三种可以添加的额外佐料也实现了饮料的抽象
// 但是它们不能单独售卖,必须在基础饮料的基础上出售
type Pearl struct {
	D Drinks
}

func (p Pearl) GetDrinksName() string {
	return "珍珠" + p.D.GetDrinksName()
}
func (p Pearl) CalcPrice() float64 {
	return 1 + p.D.CalcPrice()
}

type Pudding struct {
	D Drinks
}

func (p Pudding) GetDrinksName() string {
	return "布丁" + p.D.GetDrinksName()
}

func (p Pudding) CalcPrice() float64 {
	return 1 + p.D.CalcPrice()
}

type OldDryMom struct {
	D Drinks
}

func (o OldDryMom) GetDrinksName() string {
	return "老干妈" + o.D.GetDrinksName()
}

func (o OldDryMom) CalcPrice() float64 {
	return 1 + o.D.CalcPrice()
}