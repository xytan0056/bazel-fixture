// Package handlers implements request handlers wired to store and config.
package handlers

import (
	"strings"

	fxerrors "github.com/xytan0056/bazel-fixture/pkg/errors"
	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/pkg/strutil"
	"github.com/xytan0056/bazel-fixture/service/config"
	"github.com/xytan0056/bazel-fixture/service/store"
)

type Handlers struct {
	cfg   config.Config
	store *store.Store
	log   *logger.Logger
}

func New(cfg config.Config, s *store.Store, log *logger.Logger) *Handlers {
	return &Handlers{cfg: cfg, store: s, log: log}
}

// Echo returns the sanitized, title-cased echo of message, honoring the max
// echo length derived from MaxConnections.
func (h *Handlers) Echo(message string) (string, error) {
	if !h.cfg.Features.Echo {
		return "", fxerrors.New(fxerrors.CodeInvalidInput, "echo feature disabled")
	}
	limit := h.cfg.MaxConnections
	if limit < 1 {
		limit = 1
	}
	if limit > 4096 {
		limit = 4096
	}
	sanitized := strutil.Sanitize(message)
	titled := strutil.TitleWords(sanitized)
	if len(titled) > limit {
		titled = titled[:limit]
	}
	h.log.Info("echo " + titled)
	return titled, nil
}

// Put stores a record under id.
func (h *Handlers) Put(id, payload string) error {
	if !h.cfg.Features.Store {
		return fxerrors.New(fxerrors.CodeInvalidInput, "store feature disabled")
	}
	return h.store.Put(store.Record{ID: strings.TrimSpace(id), Payload: []byte(payload)})
}
