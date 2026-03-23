package cmd

import "github.com/BurntSushi/toml"

type ContextConfig struct {
	Context struct {
		Name string
		Root string
	}
	Env   map[string]string
	Panes []Pane
}

type Pane struct {
	Path    string
	Command string
	Split   string
}

func loadContext(path string) (ContextConfig, error) {
	var config ContextConfig
	_, err := toml.DecodeFile(path, &config)
	return config, err
}
