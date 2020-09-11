package utils

import (
	"time"
)

type Ticker interface {
	Duration() time.Duration
	Tick()
	Stop()
}

type ticker struct {
	*time.Ticker
	d time.Duration
}

func (t *ticker) Tick()                   { <-t.C }
func (t *ticker) Duration() time.Duration { return t.d }

func NewTicker(d time.Duration) Ticker {
	return &ticker{time.NewTicker(d), d}
}

type TickFunc func(d time.Duration)

func Countdown(ticker Ticker, duration time.Duration) chan time.Duration {
	remainingCh := make(chan time.Duration, 1)
	go func(ticker Ticker, dur time.Duration, remainingCh chan time.Duration) {
		for remaining := duration; remaining >= 0; remaining -= ticker.Duration() {
			remainingCh <- remaining
			ticker.Tick()
		}
		ticker.Stop()
		close(remainingCh)
	}(ticker, duration, remainingCh)
	return remainingCh
}
