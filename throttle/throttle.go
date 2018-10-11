package throttle

import (
	"sync"
	"time"
)

type Throttle interface {
	Do(key string, interval time.Duration, f func())
}

func New(options ...Option) Throttle {
	return &throttle{}
}

type throttle struct {
	doing     sync.Map
	needsNext sync.Map
}

func (t *throttle) Do(key string, interval time.Duration, f func()) {
	if _, doing := t.doing.LoadOrStore(key, nil); doing {
		t.needsNext.LoadOrStore(key, nil)
		return
	}
	go func() {
		f()
		time.Sleep(interval)
		t.doing.Delete(key)
		if _, needsNext := t.needsNext.Load(key); needsNext {
			t.needsNext.Delete(key)
			t.Do(key, interval, f)
		}
	}()
}
