package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/icrowley/fake"
)

func main() {
	m := model{}
	m.layout = "us"
	generateList(m.layout)
	m.sentence = fake.Sentence()

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("WHAAAAAT ITS BROKEN ALREAAADY ???\ndetails: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
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
				m.sentence = fake.Sentence()
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string
	if utf8.RuneCountInString(m.sentence) > 39 {
		s += fmt.Sprintf("\n     %s\n", string([]rune(m.sentence)[:39]))
	} else {
		s += fmt.Sprintf("\n     %s\n", m.sentence)
	}

	for _, item := range layouts[m.layout] {
		for _, shiftedKey := range item.sKeys {
			if shiftedKey == m.selected {
				m.shifted = true
				break
			}
		}
	}

	for _, v := range layouts[m.layout] {
		// prefix
		s += v.prefix

		var rangedSlice *[]rune

		if m.shifted {
			rangedSlice = &v.sKeys
		} else {
			rangedSlice = &v.keys
		}

		// keys
		for _, k := range *rangedSlice {
			isItClicked := m.selected == k

			if isItClicked {
				if k == m.requested {
					s += fmt.Sprintf("%s%c%s  ", colorCorrect, k, colorReset)
				} else {
					s += fmt.Sprintf("%s%c%s  ", colorWrong, k, colorReset)
				}
			} else {
				s += fmt.Sprintf("%c  ", k)
			}
		}

		// postfix
		s += v.postfix
	}
	// space
	if m.selected == ' ' {
		s += "\n\t    ┌────────────────────────┐"
		s += "\n\t    └────────────────────────┘"
	} else {
		s += "\n\t    🬞🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬏"
		s += "\n\t    🬁🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬀"
	}

	return s
}
