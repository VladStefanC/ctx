package cmd

import (
	"github.com/BurntSushi/toml"
)

type Context struct {
	Name string `toml:"name"`
	Root string `toml:"root"`
}

type ContextConfig struct {
	Context Context           `toml:"context"`
	Env     map[string]string `toml:"env"`
	Panes   []Pane            `toml:"panes"`
}

type Pane struct {
	Path    string `toml:"path"`
	Command string `toml:"command"`
	Split   string `toml:"split"`
}

func loadContext(path string) (ContextConfig, error) {
	var config ContextConfig
	_, err := toml.DecodeFile(path, &config)
	return config, err
}
