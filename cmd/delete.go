package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "This deletes an exisiting context",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		contexts, err := getContextNames()
		if err != nil {
			os.Exit(1)
		}
		var selected string
		var options []huh.Option[string]
		for _, val := range contexts {
			options = append(options, huh.NewOption(val, val))
		}
		err = huh.NewSelect[string]().Title("Pick a context you want to delete:").Options(options...).Value(&selected).Run()
		if err != nil {
			os.Exit(1)
		}
		homeDir, err := os.UserHomeDir()
		if err != nil {
			os.Exit(1)
		}
		filePath := filepath.Join(homeDir, ".config", "ctx", "contexts", selected+".toml")

		err = os.Remove(filePath)
		if err != nil {
			fmt.Errorf("Error deleting file", err)
			return
		}
		fmt.Printf("Deleted: %s\n", selected)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
