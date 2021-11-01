package proxy_test

import (
	"design/structure/proxy"
	"testing"
)

// 创建代理,并持有真实业务对象
func TestProxy(t *testing.T) {
	proxy := proxy.MakeFoodImplProxy{
		Impl: proxy.MakeFoodImpl{},
	}
	proxy.MakeChicken()
}
