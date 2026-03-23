package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", nil
		}
		return filepath.Join(home, path[2:]), nil
	}
	return path, nil
}

func sessionExists(name string) bool {
	err := exec.Command("tmux", "has-session", "-t", name).Run()
	return err == nil
}

func createSession(name, root string, panes []Pane) error {
	if err := exec.Command("tmux", "new-session", "-d", "-s", name, "-c", root).Run(); err != nil {
		return err
	}

	for i, pane := range panes {
		path, err := expandPath(pane.Path)
		if err != nil {
			return err
		}

		split := pane.Split
		if split == "" {
			split = "h"
		}
		if i == 0 {
			if pane.Command != "" {
				exec.Command("tmux", "send-keys", "-t", fmt.Sprintf("%s:0.0", name), pane.Command, "Enter").Run()
			}
			continue
		}
		if i > 1 {
			exec.Command("tmux", "select-pane", "-t", fmt.Sprintf("%s:0.%d", name, i-1)).Run()
		}
		exec.Command("tmux", "split-window", "-"+split, "-t", name, "-c", path).Run()

		if pane.Command != "" {
			exec.Command("tmux", "send-keys", "-t", fmt.Sprintf("%s:0.%d", name, i), pane.Command, "Enter").Run()
		}
	}
	return nil
}

func attachSession(name string) error {
	cmd := exec.Command("tmux", "new-session", "-t", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
