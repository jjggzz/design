package strategy_test

import (
	"design/behavior/strategy"
	"testing"
)

// 现在可以任意替换旅行的实现而不用修改Person的行为
func TestStarategy(t *testing.T) {
	var p = strategy.Person{
		T: strategy.Train{},
	}
	p.ToRun()
}
