package tui

import (
	"github.com/andrewvota/scanln/internal/tui/typing"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Quit):
			return m, tea.Quit
		}
	}

	// Delegate updates to components
	component, cmd := m.TypingComponent.Update(msg)
	m.TypingComponent = component.(*typing.Model)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
