package typing

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

const VALID_RUNES = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,.?!-_;:'\"[]()@#$%^&*+=<>/\\`~{}| "

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Backspace):
			if m.CursorPosition > 0 {
				m.Input = append(m.Input[:m.CursorPosition-1], m.Input[m.CursorPosition:]...)
				m.CursorPosition--
			}

		case strings.Contains(VALID_RUNES, msg.String()):
			runeArray := []rune(msg.String())
			m.Input = append(m.Input, runeArray...)
			m.CursorPosition += len(runeArray)
		}
	}

	cmd = m.Cursor.Focus()
	cmds = append(cmds, cmd)

	m.Cursor, cmd = m.Cursor.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
