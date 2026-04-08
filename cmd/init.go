package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializing a new TOML file for context",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, _ []string) {
		path, _ := os.Getwd()
		var root string
		var ctxName string
		var addMore bool
		var addPanes bool
		var addService bool
		envVars := make(map[string]string)
		panes := []Pane{}
		services := []Service{}

		huh.NewInput().Title("Initializing context setup").Placeholder("*context name*").Value(&ctxName).Run()
		huh.NewInput().Title("Root directory: ").Placeholder(path).Value(&root).Run()
		for {
			huh.NewConfirm().Title("Add env var?").Value(&addMore).Run()
			if !addMore {
				break
			}

			var key string
			var value string

			huh.NewInput().Title("Key").Value(&key).Run()
			huh.NewInput().Title("Value").Value(&value).Run()

			envVars[key] = value
		}

		for {
			var panePath string
			var command string
			var split string
			huh.NewConfirm().Title("Do you wanna add custom Panes?(or stick by default[split vertical and horizontal])").Value(&addPanes).Run()
			if !addPanes {
				break
			}

			huh.NewInput().Title("Add Pane path").Placeholder("~/Projects/...").Value(&panePath).Run()
			huh.NewInput().Title("Add Command").Placeholder("nvim .").Value(&command).Run()
			huh.NewSelect[string]().Title("Split Direction").Options(
				huh.NewOption("Horizontal(side by side)", "h"),
				huh.NewOption("Vertical(stacked)", "v"),
			).Value(&split).Run()

			panes = append(panes, Pane{
				Path:    panePath,
				Command: command,
				Split:   split,
			})

		}

		for {

			var serviceName string
			var service string
			huh.NewConfirm().Title("Do you want to add services you wanna check ?").Value(&addService).Run()
			if !addService {
				break
			}
			huh.NewInput().Title("Add Service name ").Placeholder("Postgres check").Value(&serviceName).Run()
			huh.NewInput().Title("Add service command").Placeholder("ex: is_psqlready").Value(&service).Run()
			services = append(services, Service{
				Name:  serviceName,
				Check: service,
			})

		}
		config := ContextConfig{
			Context: Context{
				Name: ctxName,
				Root: root,
			},
			Env:      envVars,
			Panes:    panes,
			Services: services,
		}

		home, _ := os.UserHomeDir()
		configdir := filepath.Join(home, ".config", "ctx", "contexts")
		filePath := filepath.Join(configdir, ctxName+".toml")

		f, err := os.Create(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not create config file: %v\n", err)
			os.Exit(1)
		}

		defer f.Close()

		if err := toml.NewEncoder(f).Encode(config); err != nil {
			fmt.Fprintf(os.Stderr, "could not write config file: %v\n", err)
		}

		fmt.Printf("✓ Created %s\n", filePath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
