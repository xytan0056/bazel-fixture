// Command client is a stub client that exercises the api facade in-process.
package main

import (
	"fmt"
	"os"

	"github.com/xytan0056/bazel-fixture/pkg/logger"
	"github.com/xytan0056/bazel-fixture/service/api"
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
	a := api.New(handlers.New(cfg, store.New(log), log), log)
	for _, msg := range os.Args[1:] {
		out, err := a.Serve("echo", msg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(out)
	}
}
