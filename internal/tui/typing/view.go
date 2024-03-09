package typing

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m *Model) View() string {
	styles := DefaultStyle()
	var builder strings.Builder // Declare the StringBuilder
	linesRendered := 0

	for i, r := range m.Content {
		if r == '\n' {
			linesRendered++
			if linesRendered >= 3 {
				break
			}
		}

		if i == m.CursorPosition {
			if string(r) == "\n" {
				m.CursorPosition++
			}
			m.Cursor.SetChar(string(r))
			builder.WriteString(m.Cursor.View())
		} else {
			if len(m.Input) > i && m.Input[i] == r {
				builder.WriteString(styles.CorrectText.Render(string(r)))
			} else if len(m.Input) > i && m.Input[i] != r {
				builder.WriteString(styles.IncorrectText.Render(string(r)))
			} else {
				builder.WriteString(styles.GhostText.Render(string(r)))
			}
		}
	}

	return lipgloss.Place(m.WindowWidth, m.WindowHeight, lipgloss.Center, lipgloss.Center, builder.String())
}

type Style struct {
	GhostText     lipgloss.Style
	CorrectText   lipgloss.Style
	IncorrectText lipgloss.Style
}

func DefaultStyle() Style {
	const GHOST_TEXT_COLOR = "#A1A1A1"
	const CORRECT_TEXT_COLOR = "#00FF00"
	const INCORRECT_TEXT_COLOR = "#F44336"

	return Style{
		GhostText:     lipgloss.NewStyle().Foreground(lipgloss.Color(GHOST_TEXT_COLOR)),
		CorrectText:   lipgloss.NewStyle().Foreground(lipgloss.Color(CORRECT_TEXT_COLOR)),
		IncorrectText: lipgloss.NewStyle().Foreground(lipgloss.Color(INCORRECT_TEXT_COLOR)),
	}
}
