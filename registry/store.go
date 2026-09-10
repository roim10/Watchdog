package registry

import (
	"sync"

	"github.com/roim10/Watchdog/sources"
)

type Registry struct {
	mu   sync.RWMutex
	data map[string]sources.Result
}

