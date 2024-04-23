// Package cmd implements handling of command and execution of ui
package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
	"github.com/vaclav-dvorak/goit/pkg/ui"
)

// Execute - initialize ui model
func Execute() {
	m := ui.NewModel()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		log.Fatalf("error running tea: %s", err)
	}
}
