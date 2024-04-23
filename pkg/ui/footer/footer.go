// Package footer provide footer component for ui model
package footer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/vaclav-dvorak/goit/pkg/ui/keys"
)

// Model define structure of footer component model
type Model struct {
	help    help.Model
	ShowAll bool
}

var (
	//FooterHeight - used to calculate ui height
	FooterHeight = 1
	//ExpandedHelpHeight - used to calculate ui height
	ExpandedHelpHeight = 2
)

// NewModel - returns new instance of footer component
func NewModel() Model {
	helpM := help.New()
	helpM.ShowAll = false
	return Model{
		help:    helpM,
		ShowAll: false,
	}
}

// Update - handles updates of footer component
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Help):
			m.ShowAll = !m.ShowAll
			m.help.ShowAll = m.ShowAll
		}
	}

	return m, nil
}

// View - handles view of footer component
func (m Model) View() string {
	return m.help.View(keys.Keys)
}

// SetWidth - allows to set width of component
func (m *Model) SetWidth(width int) {
	m.help.Width = width
}
