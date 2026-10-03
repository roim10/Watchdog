package sourcestore

import (
	"fmt"

	read "github.com/roim10/Watchdog/Read"
)

func NewSourceStore() *SourceStore {
	return &SourceStore{
		data: make(map[string]read.Source),
	}
}
func (s *SourceStore) Add(rs read.Source) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.data[rs.Name]
	if ok {
		return fmt.Errorf("source %q already exists", rs.Name)
	}
	s.data[rs.Name] = rs
	return nil
}
func (s *SourceStore) Delete(name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.data[name]
	if !ok {
		return fmt.Errorf("source %q does not exists", name)
	}
	delete(s.data, name)
	return nil
}
func (s *SourceStore) List() []read.Source {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sliceData := make([]read.Source, 0, len(s.data))
	for _, value := range s.data {
		sliceData = append(sliceData, value)
	}
	return sliceData
}
