package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	game   *Game
	cursor int // which to-do list item our cursor is pointing at
}

func initialModel() model {
	game := &Game{}
	game.init()
	return model{
		game:   game,
		cursor: 3,
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
			if m.cursor > 0 {
				m.cursor--
			}
		case "right", "l":
			if m.cursor < len(m.game.board)-1 {
				m.cursor++
			}
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			m.game.placePiece(m.cursor)
		}
	}
	isOver, player := m.game.isOver()
	if isOver {
		fmt.Println("\nPlayer", player, "Wins!")
		return m, tea.Quit
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	s := "For in a row!\n\n"
	cursor := "v"
	cursorRow := ""
	for i := 0; i < m.cursor; i++ {
		cursorRow += "  "
	}
	cursorRow += cursor
	s += fmt.Sprintf("%s\n", cursorRow)

	s += m.game.gameStateToString()

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
