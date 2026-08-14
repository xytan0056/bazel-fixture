// Command client is a stub client that exercises the api facade in-process.
package main

import (
	"fmt"
	"os"

	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/service/api"
	"github.com/xytan0056/bazel-fixture/service/audit"
	"github.com/xytan0056/bazel-fixture/service/config"
	"github.com/xytan0056/bazel-fixture/service/handlers"
	"github.com/xytan0056/bazel-fixture/service/store"
)

func main() {
	log := logger.New()
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	rec := audit.NewRecorder(log)
	a := api.New(handlers.New(cfg, store.New(log), rec, log), log)
	for _, msg := range os.Args[1:] {
		out, err := a.Serve("cli", "echo", msg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(out)
	}
}
