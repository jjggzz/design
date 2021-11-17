package chain_test

import (
	"design/behavior/chain"
	"testing"
)

// 创建多个处理节点,并将节点串成链(此处只有两个)
// 只需调用链首的Do方法即可执行整个链的处理方法
func TestChain(t *testing.T) {
	length := chain.NewChainNode(new(chain.LengthValidator))
	start := chain.NewChainNode(new(chain.StartWithValidator))
	length.SetNext(start)

	str := "e"
	length.Do(str)

}
