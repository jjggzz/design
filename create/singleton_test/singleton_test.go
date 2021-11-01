package singleton_test

import (
	"design/create/singleton"
	"fmt"
	"testing"
	"time"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			ins := singleton.GetInstance()
			ins.Num++
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Println(singleton.GetInstance().Num)
}
