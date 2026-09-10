package history

import (
	"github.com/roim10/Watchdog/sources"
)

func NewHistory(maxSize int) *History {
	return &History{
		data:    make(map[string][]sources.Result),
		maxSize: maxSize,
	}
}

func (h *History) Add(name string, r sources.Result) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if len(h.data[name]) >= h.maxSize {
		h.data[name] = h.data[name][1:]
	}
	h.data[name] = append(h.data[name], r)
}
func (h *History) GetHistory(name string) ([]sources.Result, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	res, ok := h.data[name]
	if !ok {
		return nil, false
	}
	cp := make([]sources.Result, len(res))
	copy(cp, res)

	return cp, true
}
