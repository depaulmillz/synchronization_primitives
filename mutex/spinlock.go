package mutex

import "sync/atomic"

// Spinlock implementation
type Spinlock struct {
	l uint64
}

// Lock locks the spinlock
func (s *Spinlock) Lock() {
	for !atomic.CompareAndSwapUint64(&s.l, 0, 1) {
	}
}

// Unlock unlocks the spinlock
func (s *Spinlock) Unlock() {
	atomic.StoreUint64(&s.l, 0)
}

// Trylock trys to lock the spinlock
// It returns whether it is sucessful
func (s *Spinlock) Trylock() bool {
	return atomic.CompareAndSwapUint64(&s.l, 0, 1)
}
