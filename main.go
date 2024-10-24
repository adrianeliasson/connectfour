package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursorPaddingLeft  string
	cursorPaddingRight string
	game               *Game
	choices            []string         // items on the to-do list
	cursor             int              // which to-do list item our cursor is pointing at
	selected           map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	game := &Game{}
	game.init()
	return model{
		cursorPaddingRight: "                     ",
		cursorPaddingLeft:  " ",
		game:               game,
		choices:            []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		selected:           make(map[int]struct{}),
	}
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			if m.cursorPaddingLeft != "" {
				m.cursorPaddingLeft = TrimSuffix(m.cursorPaddingLeft, " ")
				m.cursorPaddingRight += " "
			}
		case "right", "l":
			if m.cursorPaddingRight != "" {
				m.cursorPaddingRight = TrimSuffix(m.cursorPaddingRight, " ")
				m.cursorPaddingLeft += " "
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	topCursor := "v"
	s += fmt.Sprintf("%s%s%s\n", m.cursorPaddingLeft, topCursor, m.cursorPaddingRight)
	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

// func main() {
// 	fmt.Println("Welcome to Connect 4, the great classic game.")
// 	runConnectFour()
// }

func runConnectFour() {
	game := &Game{}
	game.init()
	game.play()
}
