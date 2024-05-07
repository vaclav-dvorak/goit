// Package ui contains ui model and its components
package ui

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"

	"github.com/vaclav-dvorak/goit/pkg/ui/footer"
	"github.com/vaclav-dvorak/goit/pkg/ui/header"
	"github.com/vaclav-dvorak/goit/pkg/ui/keys"
)

const semVerRegex string = `([0-9]+\.[0-9]+\.[0-9])`

// Model define structure of ui model
type Model struct {
	header header.Model
	footer footer.Model
	keys   keys.KeyMap
	width  int
	height int
}

// NewModel - returns new instance of ui model
func NewModel() *Model {
	return &Model{
		keys:   keys.Keys,
		header: header.NewModel(),
		footer: footer.NewModel(),
	}
}

// Init - handles initialization of model
func (m Model) Init() tea.Cmd {
	return tea.Batch(fetchGitVer, tea.EnterAltScreen)
}

// Update - handles updates of ui model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd       tea.Cmd
		headerCmd tea.Cmd
		footerCmd tea.Cmd
		cmds      []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			cmd = tea.Quit
		case key.Matches(msg, m.keys.Help):
			// probably should calculate main content height?
		}

	case tea.WindowSizeMsg:
		m.onWindowSizeChanged(msg)

	}

	m.footer, footerCmd = m.footer.Update(msg)
	m.header, headerCmd = m.header.Update(msg)

	cmds = append(
		cmds,
		cmd,
		headerCmd,
		footerCmd,
	)
	return m, tea.Batch(cmds...)
}

// View - handles view of ui
func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString(m.header.View())
	s.WriteString("\n")
	mainContent := "No sections defined..."
	s.WriteString(mainContent)
	s.WriteString("\n")
	s.WriteString(m.footer.View())
	return s.String()
}

func (m *Model) onWindowSizeChanged(msg tea.WindowSizeMsg) {
	m.footer.SetWidth(msg.Width)
	m.header.SetWidth(msg.Width)
	m.width = msg.Width
	m.height = msg.Height
	// if m.footer.ShowAll {
	// 	m.MainContentHeight = msg.Height - footer.ExpandedHelpHeight - header.HeaderHeight
	// } else {
	// 	m.MainContentHeight = msg.Height - footer.FooterHeight - header.HeaderHeight
	// }
}

func fetchGitVer() tea.Msg {
	if _, err := exec.LookPath("git"); err != nil {
		log.Fatal("we need to have git binary installed")
	}

	out, _ := exec.Command("git", "version").Output()
	re := regexp.MustCompile(semVerRegex)
	match := re.FindStringSubmatch(string(out))
	return header.FetchGitVerMsg{Version: match[0]}
}
