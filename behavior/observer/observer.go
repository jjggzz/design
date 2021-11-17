package observer

import (
	"fmt"
	"time"
)

// 定义一个主题
type Subject struct {
	obs []Observer
}

// 它有添加观察者的方法
func (sub *Subject) AddObserver(obs Observer) {
	sub.obs = append(sub.obs, obs)
}

// 也能通知观察者执行工作,这里的通知方式是异步通知,然后等待
func (sub *Subject) Notice() {
	defer func(begin time.Time) {
		fmt.Println("调用耗时:", time.Since(begin).Seconds())
	}(time.Now())
	c := make(chan struct{}, len(sub.obs))
	for _, v := range sub.obs {
		go v.update(c)
	}

	for i := 0; i < cap(c); i++ {
		<-c
	}
	close(c)
}

// 定义一个简单创建主题的方法
func NewSubject() *Subject {
	return &Subject{
		obs: make([]Observer, 0),
	}
}

// 观察者接口
type Observer interface {
	update(c chan struct{})
}

// 发邮件观察者,发邮件需要三秒钟
type EmailObs struct{}

func (EmailObs) update(c chan struct{}) {
	// time.Sleep(time.Second * 3)
	fmt.Println("发送邮件")
	c <- struct{}{}
}

// 发短信观察者,发短信需要三秒钟
type SmsObs struct{}

func (SmsObs) update(c chan struct{}) {
	// time.Sleep(time.Second * 4)
	fmt.Println("发送短信")
	c <- struct{}{}
}
