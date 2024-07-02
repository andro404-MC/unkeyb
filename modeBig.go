package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

func bigKeyb(m *model) string {
	var s, sentence string

	// Reducing or adding to the sentence to fit the box
	if utf8.RuneCountInString(m.sentence) > 61 {
		sentence += string([]rune(m.sentence)[:61])
	} else {
		sentence += m.sentence
		for i := 0; i < 61-utf8.RuneCountInString(m.sentence); i++ {
			sentence += " "
		}
	}

	// Highlighting the first letter
	sentence = colorRequested + string(
		[]rune(sentence)[:1],
	) + colorReset + string(
		[]rune(sentence)[1:],
	)

	// adding borders
	s += styleBorderNormal.Render(sentence) + "\n"

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
			s += v.prefix

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
			s += strings.TrimPrefix(v.prefix, "\n")

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

	if m.selected == ' ' {
		if m.selected == m.requested {
			s += styleBorderCorrect.MarginLeft(19).Render("                     ")
		} else {
			s += styleBorderWrong.MarginLeft(19).Render("                     ")
		}
	} else {
		s += styleBorderNormal.MarginLeft(19).Render("                     ")
	}

	styleBody := lipgloss.NewStyle().
		PaddingLeft((m.termWidh-66)/2 + 1).
		PaddingTop((m.termHeight - strings.Count(s, "\n")) / 2)

	return styleBody.Render(s)
}
