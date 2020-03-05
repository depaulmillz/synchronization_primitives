package mutex

import "testing"

func TestSpinlock(t *testing.T) {
	var s Spinlock
	if !s.Trylock() {
		t.Errorf("Unable to aquire lock although should be able to")
	}
	for i := 0; i < 100; i++ {
		if s.Trylock() {
			t.Errorf("Aquired lock although it should be locked by calling Trylock")
		}
	}
	s.Unlock()
	s.Lock()
	if s.Trylock() {
		t.Errorf("Aquired lock although it should be locked by calling Lock")
	}
}
