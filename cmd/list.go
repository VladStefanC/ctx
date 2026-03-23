package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available project contexts",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Printf("Listing all available contexts:\n")
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not find home directory", err)
			os.Exit(1)
		}
		contextsDir := filepath.Join(home, ".config", "ctx", "contexts")

		entries, err := os.ReadDir(contextsDir)
		if err != nil {
			fmt.Fprintln(os.Stderr, "could not read contexts directory", err)
		}

		for _, entry := range entries {
			if filepath.Ext(entry.Name()) == ".toml" {
				name := strings.TrimSuffix(entry.Name(), ".toml")
				fmt.Println(name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
