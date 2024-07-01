package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"gokeyb/generator"
)

func main() {
	generator.Load()

	m := model{}
	m.layout = "gb"
	generateList(m.layout)

	m.sentence = generator.Sentence()

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("WHAAAAAT ITS BROKEN ALREAAADY ???\ndetails: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape, tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyRunes, tea.KeySpace:
			m.selected = msg.Runes[0]

			if m.selected == []rune(m.sentence)[0] {
				m.requested = []rune(m.sentence)[0]
				m.sentence = strings.TrimPrefix(m.sentence, string([]rune(m.sentence)[0]))
			}

			if m.sentence == "" {
				m.sentence = generator.Sentence()
			}
		}

	case tea.WindowSizeMsg:
		m.termWidh = msg.Width
	}

	return m, nil
}

func (m model) View() string {
	return bigKeyb(&m)
}
