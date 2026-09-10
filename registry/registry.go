package registry

import (
	"github.com/roim10/Watchdog/sources"
)
	
func New() *Registry {
	return &Registry{
		data: make(map[string]sources.Result),
	}
}

func (reg *Registry) Set(name string, res sources.Result) {
	reg.mu.Lock()
	defer reg.mu.Unlock()
	reg.data[name] = res
}

func (reg *Registry) Get(name string) (sources.Result, bool) {
	reg.mu.RLock()
	defer reg.mu.RUnlock()
	res, ok := reg.data[name]
	return res, ok
}

func (reg *Registry) GetAll() map[string]sources.Result {
	reg.mu.RLock()
	defer reg.mu.RUnlock()
	copyData := make(map[string]sources.Result, len(reg.data))
	for k, v := range reg.data {
		copyData[k] = v
	}
	return copyData
}
