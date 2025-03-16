package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rogwilco/quartermaster/internal/pkg/tui"
)

// Model represents the application state
type Model struct {
	Count     int
	Items     []string
	TextInput textinput.Model
	Spinner   spinner.Model
	Help      help.Model
	Keys      tui.KeyMap
	ShowHelp  bool
	Loading   bool
}

// NewModel creates a new model with default values
func NewModel() Model {
	// Initialize text input
	ti := textinput.New()
	ti.Placeholder = "Enter an item"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30

	// Initialize spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))

	// Initialize help
	h := help.New()

	return Model{
		Count:     0,
		Items:     []string{},
		TextInput: ti,
		Spinner:   s,
		Help:      h,
		Keys:      tui.Keys,
		ShowHelp:  false,
		Loading:   false,
	}
}

// Init is the first function that will be called
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update is called when a message is received
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.Keys.Help):
			m.ShowHelp = !m.ShowHelp
		case key.Matches(msg, m.Keys.Up):
			m.Count++
		case key.Matches(msg, m.Keys.Down):
			m.Count--
		case key.Matches(msg, m.Keys.Add):
			if m.TextInput.Value() != "" {
				m.Items = append(m.Items, m.TextInput.Value())
				m.TextInput.Reset()
			}
		case key.Matches(msg, m.Keys.Submit):
			if m.TextInput.Value() != "" {
				m.Loading = true
				return m, tea.Batch(
					spinner.Tick,
					func() tea.Msg {
						// Simulate an operation that takes time
						return tui.LoadingFinishedMsg(m.TextInput.Value())
					},
				)
			}
		}
	case tui.LoadingFinishedMsg:
		m.Loading = false
		m.Items = append(m.Items, string(msg))
		m.TextInput.Reset()
	}

	// Handle spinner updates
	if m.Loading {
		spinnerCmd := m.Spinner.Tick
		cmds = append(cmds, spinnerCmd)
	}

	// Handle text input updates
	m.TextInput, cmd = m.TextInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders the current model to a string
func (m Model) View() string {
	var s strings.Builder

	// Title
	s.WriteString(tui.TitleStyle.Render("Quartermaster"))
	s.WriteString("\n\n")

	// Counter
	s.WriteString(fmt.Sprintf("Counter: %s\n", 
		tui.CounterStyle.Render(fmt.Sprintf("%d", m.Count))))

	// Items
	if len(m.Items) > 0 {
		s.WriteString("\nItems:\n")
		for i, item := range m.Items {
			s.WriteString(fmt.Sprintf("%d. %s\n", i+1, tui.ItemStyle.Render(item)))
		}
	}

	// Text input with spinner when loading
	s.WriteString("\n")
	if m.Loading {
		s.WriteString(fmt.Sprintf("%s Adding item...\n", m.Spinner.View()))
	} else {
		s.WriteString("Add item: " + m.TextInput.View() + "\n")
	}

	// Help
	helpView := m.Help.View(m.Keys)
	if !m.ShowHelp {
		// Just show a minimal version
		helpView = "Press ? for help, q to quit"
	}
	s.WriteString("\n" + tui.HelpStyle.Render(helpView))

	return s.String()
}
