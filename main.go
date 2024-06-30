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
	m.layout = "gb"
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
	return bigKeyb(&m)
}

func normalKeyb(m *model) string {
	var s string
	var sentence string
	if utf8.RuneCountInString(m.sentence) > 39 {
		sentence += string([]rune(m.sentence)[:39])
	} else {
		sentence += m.sentence
	}

	s += fmt.Sprintf(
		"\n     %s\n\n",
		colorRequested+string([]rune(sentence)[:1])+colorReset+string([]rune(sentence)[1:]),
	)

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
			isClicked := m.selected == k

			if isClicked {
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
		if m.selected == m.requested {
			s += fmt.Sprintf("\n\t    %s🬞🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬏%s", colorCorrect, colorReset)
			s += fmt.Sprintf("\n\t    %s🬁🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬀%s", colorCorrect, colorReset)
		} else {
			s += fmt.Sprintf("\n\t    %s🬞🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬏%s", colorWrong, colorReset)
			s += fmt.Sprintf("\n\t    %s🬁🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬀%s", colorWrong, colorReset)
		}
	} else {
		s += "\n\t    🬞🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬏"
		s += "\n\t    🬁🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬀"
	}

	return s
}

func bigKeyb(m *model) string {
	var s string

	s += "\n     ┌"

	for i := 0; i < 63; i++ {
		s += "─"
	}

	s += "┐"

	var sentence string
	if utf8.RuneCountInString(m.sentence) > 61 {
		sentence += string([]rune(m.sentence)[:61])
		sentence += " │"
	} else {
		sentence += m.sentence
		for i := 0; i < 61-utf8.RuneCountInString(m.sentence); i++ {
			sentence += " "
		}
		sentence += " │"
	}

	s += fmt.Sprintf(
		"\n     │ %s",
		colorRequested+string([]rune(sentence)[:1])+colorReset+string([]rune(sentence)[1:]),
	)

	s += "\n     └"

	for i := 0; i < 63; i++ {
		s += "─"
	}

	s += "┘\n"

	for _, item := range layouts[m.layout] {
		for _, shiftedKey := range item.sKeys {
			if shiftedKey == m.selected {
				m.shifted = true
				break
			}
		}
	}

	for _, v := range layouts[m.layout] {
		// top
		{
			// prefix
			s += v.prefix

			// keys
			for _, k := range v.keys {
				isClicked := m.selected == k

				if isClicked {
					if k == m.requested {
						s += fmt.Sprintf("%s%s%s", colorCorrect, "┌───┐", colorReset)
					} else {
						s += fmt.Sprintf("%s%s%s", colorWrong, "┌───┐", colorReset)
					}
				} else {
					s += "┌───┐"
				}
			}

			s += "\n"
		}
		// midle
		{
			// prefix
			s += strings.TrimPrefix(v.prefix, "\n")

			var rangedSlice *[]rune

			if m.shifted {
				rangedSlice = &v.sKeys
			} else {
				rangedSlice = &v.keys
			}

			// keys
			for _, k := range *rangedSlice {
				isClicked := m.selected == k

				if isClicked {
					if k == m.requested {
						s += fmt.Sprintf("%s│ %c │%s", colorCorrect, k, colorReset)
					} else {
						s += fmt.Sprintf("%s│ %c │%s", colorWrong, k, colorReset)
					}
				} else {
					s += fmt.Sprintf("│ %c │", k)
				}
			}
		}
		s += "\n"

		// bottom
		{
			// prefix
			s += v.prefix

			// keys
			for _, k := range v.keys {
				isClicked := m.selected == k

				if isClicked {
					if k == m.requested {
						s += fmt.Sprintf("%s%s%s", colorCorrect, "└───┘", colorReset)
					} else {
						s += fmt.Sprintf("%s%s%s", colorWrong, "└───┘", colorReset)
					}
				} else {
					s += "└───┘"
				}
			}

			s += "\n"
		}
	}

	// space
	if m.selected == ' ' {
		if m.selected == m.requested {
			s += fmt.Sprintf("\t\t\t%s┌───────────────────────┐%s", colorCorrect, colorReset)
			s += fmt.Sprintf("\n\t\t\t%s│                       │%s", colorCorrect, colorReset)
			s += fmt.Sprintf("\n\t\t\t%s└───────────────────────┘%s", colorCorrect, colorReset)
		} else {
			s += fmt.Sprintf("\t\t\t%s┌───────────────────────┐%s", colorWrong, colorReset)
			s += fmt.Sprintf("\n\t\t\t%s│                       │%s", colorWrong, colorReset)
			s += fmt.Sprintf("\n\t\t\t%s└───────────────────────┘%s", colorWrong, colorReset)
		}
	} else {
		s += "\t\t\t┌───────────────────────┐"
		s += "\n\t\t\t│                       │"
		s += "\n\t\t\t└───────────────────────┘"
	}

	return s
}
