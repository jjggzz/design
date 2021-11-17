package observer_test

import (
	"design/behavior/observer"
	"testing"
)

func TestObserver(t *testing.T) {
	var sub = observer.NewSubject()
	sub.AddObserver(observer.EmailObs{})
	sub.AddObserver(observer.SmsObs{})
	sub.Notice()
}
