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
			fmt.Fprint(os.Stderr, "could not open load %s context : %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("cd %s\n", config.Context.Root)
		for key, value := range config.Env {
			fmt.Printf("export %s=%s\n", key, value)
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
