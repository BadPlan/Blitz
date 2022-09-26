package timer

import (
	"sync"
	"time"
)

type Timer interface {
	Start()
	IsDone() bool
	PrevDelta() time.Duration
}

type NonBlockingTimer struct {
	mu       sync.Mutex
	done     bool
	duration time.Duration
	real     time.Duration
	finish   time.Time
}

func (t *NonBlockingTimer) background() {
	start := time.Now()
	for {
		if t.finish.Before(time.Now()) {
			t.done = true
			t.real = time.Now().Sub(start)
			return
		}
	}
}

func (t *NonBlockingTimer) IsDone() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.done
}

func (t *NonBlockingTimer) Start() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.done = false
	t.finish = time.Now().Add(t.duration)
	go t.background()
}

func (t *NonBlockingTimer) PrevDelta() time.Duration {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.real
}

func NewNonBlocking(duration time.Duration) *NonBlockingTimer {
	return &NonBlockingTimer{
		done:     false,
		duration: duration,
		real:     time.Duration(0),
		finish:   time.Now(),
	}
}
