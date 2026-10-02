package history

import (
	"time"

	"github.com/roim10/Watchdog/sources"
)

func NewHistory(maxSize int) *History {
	return &History{
		data:    make(map[string][]Entry),
		maxSize: maxSize,
	}
}

func (h *History) Add(name string, r sources.Result) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.data[name]) >= h.maxSize {
		h.data[name] = h.data[name][1:]
	}
	entry := Entry{
		Result: r,
		Time:   time.Now(),
	}
	h.data[name] = append(h.data[name], entry)
}
func (h *History) GetHistory(name string) ([]Entry, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	res, ok := h.data[name]
	if !ok {
		return nil, false
	}
	cp := make([]Entry, len(res))
	copy(cp, res)

	return cp, true
}
