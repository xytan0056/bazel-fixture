// Package handlers implements request handlers wired to store, config, and audit.
package handlers

import (
	"strings"

	fxerrors "github.com/xytan0056/bazel-fixture/pkg/errors"
	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/pkg/strutil"
	"github.com/xytan0056/bazel-fixture/service/audit"
	"github.com/xytan0056/bazel-fixture/service/config"
	"github.com/xytan0056/bazel-fixture/service/store"
)

const (
	minEchoLimit = 1
	maxEchoLimit = 4096
)

type Handlers struct {
	cfg   config.Config
	store *store.Store
	audit *audit.Recorder
	log   *logger.Logger
}

func New(cfg config.Config, s *store.Store, a *audit.Recorder, log *logger.Logger) *Handlers {
	return &Handlers{cfg: cfg, store: s, audit: a, log: log}
}

// Echo returns the sanitized, title-cased echo of message, honoring the max
// echo length derived from MaxConnections.
func (h *Handlers) Echo(actor, message string) (string, error) {
	if !h.cfg.Features.Echo {
		return "", fxerrors.New(fxerrors.CodeInvalidInput, "echo feature disabled")
	}
	limit := clamp(h.cfg.MaxConnections, minEchoLimit, maxEchoLimit)
	sanitized := strutil.Sanitize(message)
	titled := strutil.TitleWords(sanitized)
	if len(titled) > limit {
		titled = titled[:limit]
	}
	h.audit.Record(actor, "echo")
	h.log.Infof("echo actor=%s out=%s", actor, titled)
	return titled, nil
}

// Put stores a record under id.
func (h *Handlers) Put(actor, id, payload string) error {
	if !h.cfg.Features.Store {
		return fxerrors.New(fxerrors.CodeInvalidInput, "store feature disabled")
	}
	h.audit.Record(actor, "put")
	return h.store.Put(store.Record{ID: strings.TrimSpace(id), Payload: []byte(payload)})
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
