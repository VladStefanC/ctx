package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch <name>",
	Short: "Switch to a project context",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprint(os.Stderr, "could not find home directory", err)
			os.Exit(1)
		}

		configPath := filepath.Join(home, ".config", "ctx", "contexts", name+".toml")

		config, err := loadContext(configPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open %s context: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("cd %s\n", config.Context.Root)
		for key, value := range config.Env {
			fmt.Printf("export %s=%s\n", key, value)
		}
		if !sessionExists(name) {
			root, err := expandPath(config.Context.Root)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not expand path:", err)
				os.Exit(1)
			}
			if err := createSession(name, root, config.Panes); err != nil {
				fmt.Fprintln(os.Stderr, "could not create tmux session", err)
				os.Exit(1)
			}
		}
		if err := attachSession(name); err != nil {
			fmt.Fprintln(os.Stderr, "could not attach tmux session:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
