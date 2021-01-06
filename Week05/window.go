package Week05

import (
	"sync"
	"time"
)

type SlidingWindow struct {
	Mutex  *sync.RWMutex
	Bucket map[int64]*bucket
}

type bucket struct {
	Value float64
}

func NewSlidingWindow() *SlidingWindow {
	return &SlidingWindow{
		Bucket: make(map[int64]*bucket),
		Mutex:  &sync.RWMutex{},
	}
}

func (w *SlidingWindow) getCurrentBucket() *bucket {
	now := time.Now().Unix()
	var b *bucket
	var ok bool
	if b, ok = w.Bucket[now]; !ok {
		b = &bucket{}
		w.Bucket[now] = b
	}
	return b
}

func (w *SlidingWindow) removeOldBucket() {
	now := time.Now().Unix() - 10
	for timestamp := range w.Bucket {
		if timestamp < now {
			delete(w.Bucket, timestamp)
		}
	}
}

func (w *SlidingWindow) Increment(i float64) {
	if i == 0 {
		return
	}
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	b := w.getCurrentBucket()
	b.Value += i
	w.removeOldBucket()
}

func (w *SlidingWindow) UpdateMax(max float64) {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	b := w.getCurrentBucket()

	if max > b.Value {
		b.Value = max
	}
	w.removeOldBucket()
}

func (w *SlidingWindow) Sum(now time.Time) (sum float64) {
	w.Mutex.RLock()
	defer w.Mutex.RUnlock()

	for timestamp, b := range w.Bucket {
		if timestamp >= now.Unix()-10 {
			sum += b.Value
		}
	}

	return
}

func (w *SlidingWindow) Max() (max float64) {
	w.Mutex.RLock()
	defer w.Mutex.RUnlock()

	for timestamp, b := range w.Bucket {
		if timestamp >= time.Now().Unix()-10 && b.Value > max {
			max = b.Value
		}
	}
	return
}

func (w *SlidingWindow) Avg(now time.Time) float64 {
	return w.Sum(now) / 10
}
