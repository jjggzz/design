package chain

import (
	"fmt"
	"strings"
)

// 如果改为首字母大写则其在其他包中也能实现它,然后自定义行为
// 定义责任链上的处理行为抽象
type validator interface {
	Do(str string) bool
}

// 责任链节点的抽象,可以设置下一节点或者执行责任链的行为
type ChainNode interface {
	SetNext(next ChainNode)
	Do(str string)
}

// 暴露出的创建节点的函数,在内部实例化了被隐藏起来的细节
func NewChainNode(v validator) ChainNode {
	return &chainNode{
		validator: v,
	}
}

// 责任链上的节点实现,其中包含了行为抽象和下一节点(如果有)的引用
type chainNode struct {
	validator
	next ChainNode
}

// 为当前责任节点设置下一节点
func (a *chainNode) SetNext(next ChainNode) {
	a.next = next
}

// 节点的执行逻辑是
// 执行当前节点的处理行为.如果有下一节点则执行下一节点的Do操作
func (a *chainNode) Do(str string) {
	// 调用当前节点的处理逻辑,如果返回true,并且存在下一节点则调用下一节点
	if a.validator.Do(str) && nil != a.next {
		a.next.Do(str)
	}
}

// 定义责任链处理行为的具体实现
type LengthValidator struct{}

func (l *LengthValidator) Do(str string) bool {
	if len(str) <= 0 {
		fmt.Println("str len <= 0", str)
		return false
	}
	return true
}

// 定义责任链处理行为的具体实现
type StartWithValidator struct{}

func (s *StartWithValidator) Do(str string) bool {
	if !strings.HasPrefix(str, "H") {
		fmt.Println("str prefix is not H", str)
		return false
	}
	return true
}
