package mutex

// Trylocker allows for locking, unlocking, and trylocking
type Trylocker interface {
	Lock()
	Unlock()
	Trylock() bool
}
