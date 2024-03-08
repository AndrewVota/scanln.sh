package tui

import (
	"github.com/andrewvota/scanln/internal/tui/typing"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	// Global
	WindowWidth  int
	WindowHeight int
	Keys         KeyMap

	// Components & Properties
	TypingComponent *typing.Model
}

func New() *Model {
	return &Model{
		WindowWidth:  0,
		WindowHeight: 0,
		Keys:         DefaultKeyMap(),

		TypingComponent: typing.New(),
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(m.TypingComponent.Init())
}

type KeyMap struct {
	Quit key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "esc"),
		),
	}
}
