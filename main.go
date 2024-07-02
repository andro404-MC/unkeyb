package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

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
		m.termWidth = msg.Width
		m.termHeight = msg.Height
	}

	return m, nil
}

func (m model) View() string {
	// Getting the big template (more comming soon)
	visual := bigKeyb(&m)

	// Getting visual height & width
	visualWidth := lipgloss.Width(visual)
	visualHeight := lipgloss.Height(visual)

	// Check if there is enough space
	if m.termWidth < visualWidth || m.termHeight < visualHeight {
		visual = "Terminal size too small:\n"

		// Coloring
		// Width
		if m.termWidth < visualWidth {
			visual += fmt.Sprintf("Width = %s%d%s",
				generator.AnsiToString(1), m.termWidth, generator.AnsiReset)
		} else {
			visual += fmt.Sprintf("Width = %s%d%s",
				generator.AnsiToString(2), m.termWidth, generator.AnsiReset)
		}

		// Height
		if m.termHeight < visualHeight {
			visual += fmt.Sprintf(" Height = %s%d%s\n\n",
				generator.AnsiToString(1), m.termHeight, generator.AnsiReset)
		} else {
			visual += fmt.Sprintf(" Height = %s%d%s\n\n",
				generator.AnsiToString(2), m.termHeight, generator.AnsiReset)
		}

		visual += "Needed for current config:\n"
		visual += fmt.Sprintf("Width = %d Height = %d", visualWidth, visualHeight)
	}

	visual = lipgloss.Place(m.termWidth, m.termHeight, lipgloss.Center, lipgloss.Center, visual)
	return visual
}
