package adapter_test

import (
	"design/structure/adapter"
	"testing"
)

func TestAdapter(t *testing.T) {
	manager := adapter.ShopManagerAdapter{
		Salesclerk: adapter.SalesclerkImpl{},
	}
	manager.Manager()
	manager.CalcPrice()
}
