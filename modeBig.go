package main

import (
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"

	"gokeyb/generator"
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
	sentence = styleRequested.Render(string([]rune(sentence)[:1])) + string([]rune(sentence)[1:])

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

	for _, r := range layouts[m.layout] {
		var rangedSlice *[]rune
		var rowStrings []string

		if m.shifted {
			rangedSlice = &r.sKeys
		} else {
			rangedSlice = &r.keys
		}

		for _, k := range *rangedSlice {
			isClicked := m.selected == k
			if isClicked {
				if k == m.requested {
					rowStrings = append(rowStrings, styleBorderCorrect.Render(string(k)))
				} else {
					rowStrings = append(rowStrings, styleBorderWrong.Render(string(k)))
				}
			} else {
				rowStrings = append(rowStrings, styleBorderNormal.Render(string(k)))
			}
		}

		s += lipgloss.NewStyle().MarginLeft(r.prefix).
			Render(lipgloss.JoinHorizontal(lipgloss.Right, rowStrings...))
		s += "\n"
	}

	// Space bar
	if m.selected == ' ' {
		if m.selected == m.requested {
			s += styleBorderCorrect.MarginLeft(19).Render(generator.Spaces(21))
		} else {
			s += styleBorderWrong.MarginLeft(19).Render(generator.Spaces(21))
		}
	} else {
		s += styleBorderNormal.MarginLeft(19).Render(generator.Spaces(21))
	}

	// Defining the body style with centering
	styleBody := lipgloss.NewStyle().
		PaddingLeft((m.termWidh-66)/2 + 1).
		PaddingTop((m.termHeight - strings.Count(s, "\n")) / 2)

	return styleBody.Render(s)
}
