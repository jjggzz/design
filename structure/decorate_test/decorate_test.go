package decorate_test

import (
	"design/structure/decorate"
	"fmt"
	"testing"
)

// 现在我需要一份加老干妈,两份珍珠,布丁的椰子汁

// 装饰者模式除了能为被装饰对象添加额外的功能外,还可以解决
// 像这种恶心的组合嵌套问题,避免了类爆炸的问题(排列组合太多了)
func TestDecorate(t *testing.T) {
	// 首先有一个椰子汁
	coc := decorate.Coconut{}
	// 包装一份布丁
	pud := decorate.Pudding{
		D: coc,
	}
	// 包装一份珍珠
	pearl1 := decorate.Pearl{
		D: pud,
	}
	// 包装一份珍珠
	pearl2 := decorate.Pearl{
		D: pearl1,
	}
	// 包装一份老干妈
	drink := decorate.OldDryMom{
		D: pearl2,
	}

	fmt.Println(drink.GetDrinksName())
	fmt.Println(drink.CalcPrice())
}
