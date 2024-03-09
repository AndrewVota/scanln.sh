package tui

func (m *Model) View() string {
	return m.TypingComponent.View()
}

type Style struct{}

func DefaultStyle() Style {
	return Style{}
}
