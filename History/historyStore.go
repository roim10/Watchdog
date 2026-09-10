package history

import (
	"sync"

	"github.com/roim10/Watchdog/sources"
)

type History struct {
	mu      sync.RWMutex
	data    map[string][]sources.Result
	maxSize int
}
