// Package store is an in-memory record store used by handlers.
package store

import (
	"sync"

	fxerrors "github.com/xytan0056/bazel-fixture/pkg/errors"
	"github.com/xytan0056/bazel-fixture/pkg/logger"
)

type Record struct {
	ID      string
	Payload []byte
}

type Store struct {
	mu      sync.RWMutex
	records map[string]Record
	log     *logger.Logger
}

func New(log *logger.Logger) *Store {
	return &Store{records: make(map[string]Record), log: log}
}

func (s *Store) Put(r Record) error {
	if r.ID == "" {
		return fxerrors.New(fxerrors.CodeInvalidInput, "record id is required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.records[r.ID] = r
	s.log.Info("store put " + r.ID)
	return nil
}

func (s *Store) Get(id string) (Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	r, ok := s.records[id]
	if !ok {
		return Record{}, fxerrors.New(fxerrors.CodeNotFound, "record "+id)
	}
	return r, nil
}

func (s *Store) List() []Record {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Record, 0, len(s.records))
	for _, r := range s.records {
		out = append(out, r)
	}
	return out
}
