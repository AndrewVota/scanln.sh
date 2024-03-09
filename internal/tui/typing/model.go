package typing

import (
	"time"
	"unicode"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

const TIMER_DURATION = 30 * time.Second

// NOTE: This should be a network call done in Init()
// - Angron, Primarch of the World Eaters
const CONTENT = "What would you know of struggle, perfect son? When have you fought against the mutilation of your mind? When have you had to do anything other than tally compliance's and polish your armor? The people of your world named you \"Great One\". The people of mine called me slave. Which one of us landed on a paradise of civilization to be raised by a foster father, Roboute? Which one of us was given armies to lead after training in the halls of the Macraggian High Riders? Which one of us inherited a strong, cultured kingdom? And which one of us had to rise up against a kingdom with nothing but a horde of starving slaves? Which one of us was a child enslaved on a world of monsters, with his brain cut up by carving knives? Listen to your blue clad wretches yelling courage and honor, courage and honor, courage and honor! Do you even know the meaning of those words? Courage is fighting the kingdom which enslaves you, no matter that their armies outnumber yours by ten-thousand to one. You know nothing of courage! Honor is resisting a tyrant when all others suckle and grow fat on the hypocrisy he feeds them. You know nothing of honor!"

type Model struct {
	// Global
	WindowWidth  int
	WindowHeight int
	Keys         KeyMap
	Active       bool

	// Components & Properties
	Timer          timer.Model
	Cursor         cursor.Model
	CursorPosition int
	Content        []rune
	Input          []rune
}

func New() *Model {
	return &Model{
		WindowWidth:  0,
		WindowHeight: 0,
		Keys:         DefaultKeyMap(),
		Active:       true,

		Timer:          timer.New(TIMER_DURATION),
		Cursor:         cursor.New(),
		CursorPosition: 0,
		Content:        []rune(InsertNewlines(CONTENT, 80)), // NOTE: This will be a file read in the `Init()` function
		Input:          []rune{},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

type KeyMap struct {
	Backspace key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Backspace: key.NewBinding(
			key.WithKeys("backspace"),
		),
	}
}

func InsertNewlines(text string, maxLineLength int) string {
	wordStartIndex := 0
	currentLineLength := 0
	result := []rune(text) // Work with a rune slice

	for i, r := range result {
		if unicode.IsSpace(r) {
			wordStartIndex = i + 1
		} else {
			currentLineLength++
		}

		if currentLineLength >= maxLineLength {
			if wordStartIndex < i {
				for j := i; j >= 0; j-- {
					if unicode.IsSpace(rune(result[j])) {
						i = j
						break
					}
				}
			}

			result = append(result[:i+1], append([]rune{'\n'}, result[i+1:]...)...)
			currentLineLength = i - wordStartIndex
			wordStartIndex = i + 2
		}
	}

	return string(result) // Convert the result back to a string
}
