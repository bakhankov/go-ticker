package ticker

import (
	"context"
	"sync"
	"time"
)

// A ManagedTicker holds a channel that delivers “ticks” of a clock
// at intervals.
type ManagedTicker struct {
	C               <-chan time.Time
	c               chan time.Time
	timeTicker      *time.Ticker
	initialDuration time.Duration
	ctx             context.Context
	cancel          context.CancelFunc
	mux             sync.Mutex
}

// NewTicker returns a new ManagedTicker containing a channel that will send
// the current time on the channel after each tick. The period of the
// ticks is specified by the duration argument. The ticker will adjust
// the time interval or drop ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, NewTicker will
// panic. Stop the ticker to release associated resources.
// `tickOnInit` specifies whether the ticker should send an initial tick
// immediately upon creation, otherwise it will wait for
// the first tick interval to elapse before sending the first tick.
func NewTicker(d time.Duration, tickOnInit bool) *ManagedTicker {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan time.Time, 1)
	t := ManagedTicker{
		C:               ch,
		c:               ch,
		timeTicker:      time.NewTicker(d),
		initialDuration: d,
		ctx:             ctx,
		cancel:          cancel,
	}

	go func() {
		for {
			select {
			case <-t.ctx.Done():
				return
			case tm := <-t.timeTicker.C:
				t.tick(tm)
			}
		}
	}()

	if tickOnInit {
		t.tick(time.Now())
	}
	return &t
}

// Stop turns off a ticker. After Stop, no more ticks will be sent.
// Stop does not close the channel, to prevent a concurrent goroutine
// reading from the channel from seeing an erroneous "tick".
func (t *ManagedTicker) Stop() {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.timeTicker.Stop()
	t.cancel()
	close(t.c)
}

// Reset stops a ticker and resets its period to the specified duration.
// The next tick will arrive after the new period elapses. The duration d
// must be greater than zero; if not, Reset will panic.
func (t *ManagedTicker) Reset(d time.Duration) {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.timeTicker.Reset(d)
}

// TickAfter schedules a tick to be sent after the specified duration.
func (t *ManagedTicker) TickAfter(d time.Duration) {
	time.AfterFunc(d, t.Tick)
}

// Tick sends a tick immediately, if the ticker is still running.
// If the ticker has been stopped, it will not send a tick.
func (t *ManagedTicker) Tick() {
	t.tick(time.Now())
}

func (t *ManagedTicker) tick(tm time.Time) {
	t.mux.Lock()
	defer t.mux.Unlock()

	if t.ctx.Err() != nil {
		return
	}

	select {
	case t.c <- tm:
		t.timeTicker.Reset(t.initialDuration)
	default:
	}
}
