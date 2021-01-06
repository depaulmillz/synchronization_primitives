package mutex

import (
	"sync/atomic"
	"time"
)

// SharedMutex ...
type SharedMutex struct {
	i *int64 // if *i is -1 its exclusive locked, if it is 0 its unlocked, if it is >0 its shared locked
}

// Init ...
func (m *SharedMutex) Init() {
	m.i = new(int64)
	atomic.StoreInt64(m.i, 0)
}

// TryExLock ...
func (m *SharedMutex) TryExLock() bool {
	return atomic.CompareAndSwapInt64(m.i, 0, -1)
}

// ExLock ...
func (m *SharedMutex) ExLock() {
	for !m.TryExLock() {
	}
}

// TimeoutExLock ...
func (m *SharedMutex) TimeoutExLock(d time.Duration) bool {
	t := time.NewTimer(d)
	for !m.TryExLock() {
		select {
		case <-t.C:
			return false
		default:
		}
	}
	t.Stop()
	return true
}

// ExUnlock ...
func (m *SharedMutex) ExUnlock() {
	atomic.StoreInt64(m.i, 0)
}

// TrySharedLock ...
func (m *SharedMutex) TrySharedLock() bool {
	expected := atomic.LoadInt64(m.i)
	if expected >= 0 {
		return atomic.CompareAndSwapInt64(m.i, expected, expected+1)
	}
	return false
}

// TimeoutSharedLock ...
func (m *SharedMutex) TimeoutSharedLock(d time.Duration) bool {
	t := time.NewTimer(d)
	for !m.TrySharedLock() {
		select {
		case <-t.C:
			return false
		default:
		}
	}
	t.Stop()
	return true
}

// SharedLock ...
func (m *SharedMutex) SharedLock() {
	for !m.TrySharedLock() {
	}
}

// SharedUnlock ...
func (m *SharedMutex) SharedUnlock() {
	atomic.AddInt64(m.i, -1)
}

// Lock ...
func (m *SharedMutex) Lock() {
	m.ExLock()
}

// Unlock ...
func (m *SharedMutex) Unlock() {
	m.ExUnlock()
}

// Trylock ...
func (m *SharedMutex) Trylock() {
	m.Trylock()
}
