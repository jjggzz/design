package build_test

import (
	"design/create/build"
	"fmt"
	"testing"
)

// 通过不同的构建器构建同一种对
// 不同的构建器有不同的构建校验规则
func TestBuild(t *testing.T) {
	builder1 := new(build.YoungBuilder)
	persion1, err1 := builder1.SetName("zs").SetAge(20).SetGender("man").Build()
	fmt.Println(persion1, err1)

	builder2 := new(build.ElderlyBuilder)
	persion2, err2 := builder2.SetName("zs").SetAge(20).SetGender("man").Build()
	fmt.Println(persion2, err2)

}
