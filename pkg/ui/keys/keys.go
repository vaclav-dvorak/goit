// Package keys provide key functionality for ui model
package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

// KeyMap - strucut for keyboard bindings
type KeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Refresh key.Binding
	Help    key.Binding
	Quit    key.Binding
}

// ShortHelp - return binding for short help
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help}
}

// FullHelp - return binding for full help
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		k.navigationKeys(),
		k.appKeys(),
		k.quitAndHelpKeys(),
	}
}

func (k KeyMap) navigationKeys() []key.Binding {
	return []key.Binding{
		k.Up,
		k.Down,
	}
}

func (k KeyMap) appKeys() []key.Binding {
	return []key.Binding{
		k.Refresh,
	}
}

func (k KeyMap) quitAndHelpKeys() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// Keys - binded keys and their meaning
var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Refresh: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
