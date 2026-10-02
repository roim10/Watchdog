package history

import (
	"sync"
)

type History struct {
	mu      sync.RWMutex
	data    map[string][]Entry
	maxSize int
}
