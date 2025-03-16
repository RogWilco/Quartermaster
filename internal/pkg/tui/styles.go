package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Define styles for the application
var (
	// TitleStyle is used for the application title
	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingLeft(2).
		PaddingRight(2).
		MarginBottom(1)

	// CounterStyle is used for the counter display
	CounterStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#2D3748")).
		Padding(0, 1).
		MarginTop(1)

	// ItemStyle is used for list items
	ItemStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA"))

	// HelpStyle is used for the help text
	HelpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		MarginTop(1)
)
