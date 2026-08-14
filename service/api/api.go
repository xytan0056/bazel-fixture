// Package api is the outer service facade. It owns the request lifecycle and
// delegates to handlers.
package api

import (
	fxerrors "github.com/xytan0056/bazel-fixture/pkg/errors"
	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/service/handlers"
)

type API struct {
	h   *handlers.Handlers
	log *logger.Logger
}

func New(h *handlers.Handlers, log *logger.Logger) *API {
	return &API{h: h, log: log}
}

// Serve routes an incoming request name to the appropriate handler.
func (a *API) Serve(op string, arg string) (string, error) {
	switch op {
	case "echo":
		return a.h.Echo(arg)
	case "put":
		return "", a.h.Put(arg, arg)
	default:
		return "", fxerrors.New(fxerrors.CodeInvalidInput, "unknown op "+op)
	}
}
