// Package config loads fixture service configuration from embedded YAML and
// JSON defaults. The YAML file is the primary source of truth; the JSON file
// carries the numeric defaults.
package config

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/goccy/go-yaml"

	fxerrors "github.com/xytan0056/bazel-fixture/pkg/errors"
)

//go:embed config.yaml
var configYAML []byte

//go:embed defaults.json
var defaultsJSON []byte

type Features struct {
	Echo  bool `yaml:"echo"`
	Store bool `yaml:"store"`
}

type Config struct {
	ServiceName      string   `yaml:"service_name"`
	Environment      string   `yaml:"environment"`
	Features         Features `yaml:"features"`
	MaxConnections   int      `json:"max_connections"`
	RequestTimeoutMs int      `json:"request_timeout_ms"`
	ListenAddress    string   `json:"listen_address"`
}

// Load parses the embedded YAML/JSON and returns the merged Config.
func Load() (Config, error) {
	var c Config
	if err := yaml.Unmarshal(configYAML, &c); err != nil {
		return Config{}, fxerrors.Wrap(fxerrors.CodeInternal, "parse config.yaml", err)
	}
	if err := json.Unmarshal(defaultsJSON, &c); err != nil {
		return Config{}, fxerrors.Wrap(fxerrors.CodeInternal, "parse defaults.json", err)
	}
	if c.ListenAddress == "" {
		return Config{}, fxerrors.New(fxerrors.CodeInvalidInput, "listen_address is required")
	}
	return c, nil
}

func (c Config) String() string {
	return fmt.Sprintf("service=%s env=%s listen=%s", c.ServiceName, c.Environment, c.ListenAddress)
}
