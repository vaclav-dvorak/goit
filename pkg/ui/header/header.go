// Package header provide header component for ui model
package header

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	version   = "devel"
	logoStyle = lipgloss.NewStyle().
			PaddingRight(2).
			Foreground(lipgloss.Color("63"))

	infoStyle = lipgloss.NewStyle().
			PaddingLeft(2).
			Foreground(lipgloss.Color("63"))
	logoSlice = []string{ //ascii generator - font "Thin"
		"          o|",
		",---.,---..|---",
		"|   ||   |||",
		"`---|`---'``---'",
		"`---'",
	}
	logo = logoStyle.Render(strings.Join(logoSlice, "\n"))
)

// Model define structure of header model
type Model struct {
	version string
	gitVer  string
	width   int
}

// NewModel - returns new instance of header component
func NewModel() Model {
	return Model{
		version: version,
		gitVer:  "0.0",
	}
}

// FetchGitVerMsg - struct to send tea.Msg into header component
type FetchGitVerMsg struct {
	Version string
}

// Update - handles updates of header component
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FetchGitVerMsg:
		m.gitVer = msg.Version
	}

	return m, nil
}

// View - handles view of header component
func (m Model) View() string {
	infoBar := infoStyle.Render(fmt.Sprintf("goit v: %s\ngit v:  v%s", version, m.gitVer))
	w := m.width - lipgloss.Width(infoBar) - lipgloss.Width(logo) - 2 // 2 for two border lines
	if w < 0 {
		w = 0
	}
	spacing := lipgloss.NewStyle().
		// Background(m.ctx.Theme.SelectedBackground).
		Render(strings.Repeat(" ", w))

	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Render(
		lipgloss.JoinHorizontal(lipgloss.Center, infoBar, spacing, logo),
	)
}

// SetWidth - allows to set width of component
func (m *Model) SetWidth(width int) {
	m.width = width
}
