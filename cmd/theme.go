package cmd

import "github.com/charmbracelet/lipgloss"

var (
	Title   = lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Bold(true)
	Success = lipgloss.NewStyle().Foreground(lipgloss.Color("46"))
	Error   = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	Dim     = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	Icon    = lipgloss.NewStyle().Foreground(lipgloss.Color("226"))
)
