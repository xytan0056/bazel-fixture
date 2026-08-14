// Command server boots the fixture service.
package main

import (
	"fmt"
	"os"

	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/pkg/version"
	"github.com/xytan0056/bazel-fixture/service/api"
	"github.com/xytan0056/bazel-fixture/service/audit"
	"github.com/xytan0056/bazel-fixture/service/config"
	"github.com/xytan0056/bazel-fixture/service/handlers"
	"github.com/xytan0056/bazel-fixture/service/store"
)

func main() {
	log := logger.New()
	log.Info(version.Name() + " " + version.BuildTag())

	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	log.Info(cfg.String())

	rec := audit.NewRecorder(log)
	a := api.New(handlers.New(cfg, store.New(log), rec, log), log)
	out, err := a.Serve("system", "echo", "hello & world")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(out)
}
