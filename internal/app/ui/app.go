package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rogwilco/quartermaster/internal/app/model"
)

// Run starts the application
func Run() error {
	m := model.NewModel()
	p := tea.NewProgram(m)
	
	_, err := p.Run()
	return err
}
