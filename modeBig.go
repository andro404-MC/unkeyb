package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func bigKeyb(m *model) string {
	var s, padding string
	for i := 0; i < (m.termWidh-66)/2+1; i++ {
		padding += " "
	}

	s += "\n" + padding + "┌"

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
		"\n"+padding+"│ %s",
		colorRequested+string([]rune(sentence)[:1])+colorReset+string([]rune(sentence)[1:]),
	)

	s += "\n" + padding + "└"

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
		var rangedSlice *[]rune

		if m.shifted {
			rangedSlice = &v.sKeys
		} else {
			rangedSlice = &v.keys
		}

		// top
		{
			// prefix
			s += padding + v.prefix

			// keys
			for _, k := range *rangedSlice {
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
			s += padding + strings.TrimPrefix(v.prefix, "\n")

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
			s += padding + v.prefix

			// keys
			for _, k := range *rangedSlice {
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
			s += fmt.Sprintf(padding+"                   %s┌───────────────────────┐%s",
				colorCorrect, colorReset)

			s += fmt.Sprintf("\n"+padding+"                   %s│                       │%s",
				colorCorrect, colorReset)

			s += fmt.Sprintf("\n"+padding+"                   %s└───────────────────────┘%s",
				colorCorrect, colorReset)
		} else {
			s += fmt.Sprintf(padding+"                   %s┌───────────────────────┐%s",
				colorWrong, colorReset)

			s += fmt.Sprintf("\n"+padding+"                   %s│                       │%s",
				colorWrong, colorReset)

			s += fmt.Sprintf("\n"+padding+"                   %s└───────────────────────┘%s",
				colorWrong, colorReset)
		}
	} else {
		s += padding + "                   ┌───────────────────────┐"
		s += "\n" + padding + "                   │                       │"
		s += "\n" + padding + "                   └───────────────────────┘"
	}

	return s
}
