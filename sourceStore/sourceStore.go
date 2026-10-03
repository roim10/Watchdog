package sourcestore

import (
	"sync"

	read "github.com/roim10/Watchdog/Read"
)

type SourceStore struct {
	mu   sync.RWMutex
	data map[string]read.Source
}
