package main

import (
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type row struct {
	prefix  string
	postfix string
	keys    []rune
}

var gb = [...]row{
	{
		prefix: "\n\t", postfix: "\n\n",
		keys: []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
	},
	{
		prefix: "\t ", postfix: "\n\n",
		keys: []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']'},
	},
	{
		prefix: "\t  ", postfix: "\n\n",
		keys: []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '@', '~'},
	},
	{
		prefix: "\t ", postfix: "\n",
		keys: []rune{'\\', 'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/'},
	},
}

type model struct {
	requested rune
	next      rune
	selected  rune
}

var keyList []rune

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape:
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
	s += fmt.Sprintf("\trequested : %c\n", m.next)

	for _, v := range gb {
		// prefix
		s += v.prefix

		// keys
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

		// postfix
		s += v.postfix
	}

	return s
}

func main() {
	generateList("gb")

	m := model{}
	m.next = keyList[rand.Intn(len(keyList))]

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("WHAAAAAT ITS BROKEN ALREAAADY ???\ndetails: %v", err)
		os.Exit(1)
	}
}

func generateList(layout string) {
	if layout != "" {
		for _, v := range gb {
			keyList = append(keyList, v.keys...)
		}
	}
}
