package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

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
