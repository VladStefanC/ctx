package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

func getContextNames() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	contextDir := filepath.Join(home, ".config", "ctx", "contexts")
	entries, err := os.ReadDir(contextDir)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".toml" {
			names = append(names, strings.TrimSuffix(entry.Name(), ".toml"))
		}
	}

	return names, nil
}
