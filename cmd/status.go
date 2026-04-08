package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check status of different services running in the context",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		home, _ := os.UserHomeDir()
		configPath := filepath.Join(home, ".config", "ctx", "contexts", name+".toml")
		var wg sync.WaitGroup

		config, err := loadContext(configPath)
		fmt.Println("config data: ", config)
		if err != nil {
			fmt.Fprint(os.Stderr, "could not load context : '%s': %v\n", name, err)
			os.Exit(1)
		}

		results := make(chan ServiceResult, len(config.Services))
		for _, service := range config.Services {
			wg.Add(1)
			go func(s Service) {
				defer wg.Done()
				cmd := exec.Command("sh", "-c", s.Check)
				err := cmd.Run()
				results <- ServiceResult{s.Name, err == nil}
			}(service)
		}
		done := make(chan struct{})
		go func() {
			time.Sleep(1 * time.Second)
			dots := []string{".", "..", "...", ""}
			i := 0
			for {
				select {
				case <-done:
					return
				default:
					fmt.Printf("\rChecking services %s", dots[i%len(dots)])
					time.Sleep(500 * time.Millisecond)
					i++
				}
			}
		}()

		wg.Wait()
		close(results)
		close(done)

		fmt.Println()
		if len(config.Services) == 0 {
			fmt.Println("No services defined")
			return
		}
		var serviceResults []ServiceResult
		for r := range results {
			serviceResults = append(serviceResults, r)
		}

		for _, r := range serviceResults {
			status := "✓"
			if !r.Running {
				status = "x"
			}
			fmt.Printf("%s %s\n", status, r.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
