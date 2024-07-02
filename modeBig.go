package main

import (
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"

	"gokeyb/generator"
)

func bigKeyb(m *model) string {
	// Layer
	var layerSentence, layerKeyb, layerSpace string
	var shifted bool

	// Reducing or adding to the sentence to fit the box
	if utf8.RuneCountInString(m.sentence) > 61 {
		layerSentence = string([]rune(m.sentence)[:61])
	} else {
		layerSentence = m.sentence
		for i := 0; i < 61-utf8.RuneCountInString(m.sentence); i++ {
			layerSentence += " "
		}
	}

	// Highlighting the first letter
	layerSentence = styleRequested.Render(
		string([]rune(layerSentence)[:1])) +
		string([]rune(layerSentence)[1:])

	// Adding borders
	layerSentence = styleBorderNormal.Render(layerSentence)

	// Checking if shifted
	for _, item := range layouts[m.layout] {
		for _, shiftedKey := range item.sKeys {
			if shiftedKey == m.selected {
				shifted = true
				break
			}
		}
	}

	// Drawing Rows
	var rows []string
	for _, r := range layouts[m.layout] {
		var rangedSlice *[]rune
		var keys []string

		// Assigning appropriate slice to rangedSlice
		if shifted {
			rangedSlice = &r.sKeys
		} else {
			rangedSlice = &r.keys
		}

		// Creating keys boxes
		for _, k := range *rangedSlice {
			isClicked := m.selected == k
			if isClicked {
				if k == m.requested {
					keys = append(keys, styleBorderCorrect.Render(string(k)))
				} else {
					keys = append(keys, styleBorderWrong.Render(string(k)))
				}
			} else {
				keys = append(keys, styleBorderNormal.Render(string(k)))
			}
		}

		// Merging to row
		rows = append(rows,
			lipgloss.NewStyle().
				MarginLeft(r.prefix).
				Render(lipgloss.JoinHorizontal(lipgloss.Right, keys...)),
		)
	}

	// Mergin rows to the Keyboard layer
	layerKeyb = lipgloss.JoinVertical(lipgloss.Left, rows...)

	// Space bar
	if m.selected == ' ' {
		if m.selected == m.requested {
			layerSpace = styleBorderCorrect.MarginLeft(19).Render(generator.Spaces(21))
		} else {
			layerSpace = styleBorderWrong.MarginLeft(19).Render(generator.Spaces(21))
		}
	} else {
		layerSpace += styleBorderNormal.MarginLeft(19).Render(generator.Spaces(21))
	}

	// Merging layers
	visual := lipgloss.JoinVertical(lipgloss.Left, layerSentence, layerKeyb, layerSpace)

	// Calculating visual Height and Widh
	visualWidh := (m.termWidh-66)/2 + 1
	visualHeight := (m.termHeight - strings.Count(visual, "\n")) / 2

	// Defining the body style with centering
	styleBody := lipgloss.NewStyle().MarginLeft(visualWidh).MarginTop(visualHeight)

	return styleBody.Render(visual)
}
