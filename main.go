package main

import (
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := model{}
	m.layout = "gb"
	generateList(m.layout)
	m.next = keyList[rand.Intn(len(keyList))]

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

		case tea.KeyRunes:
			m.selected = msg.Runes[0]

			if m.selected == m.next {
				m.requested = m.next
				m.next = keyList[rand.Intn(len(keyList))]
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string
	m.shifted = false
	s += fmt.Sprintf("\n\trequested : %c\n", m.next)

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

		// keys
		if m.shifted {
			for _, k := range v.sKeys {
				isItClicked := m.selected == k

				if isItClicked {
					if k == m.requested {
						s += fmt.Sprintf("\033[38;5;27m%c\033[0m  ", k)
					} else {
						s += fmt.Sprintf("\033[38;5;196m%c\033[0m  ", k)
					}
				} else {
					s += fmt.Sprintf("%c  ", k)
				}
			}
		} else {
			for _, k := range v.keys {
				isItClicked := m.selected == k

				if isItClicked {
					if k == m.requested {
						s += fmt.Sprintf("\033[38;5;27m%c\033[0m  ", k)
					} else {
						s += fmt.Sprintf("\033[38;5;196m%c\033[0m  ", k)
					}
				} else {
					s += fmt.Sprintf("%c  ", k)
				}
			}
		}

		// postfix
		s += v.postfix
	}

	return s
}
