package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"

	"gokeyb/generator"
)

func bigKeyb(m *model) string {
	// Layers
	var layerSentence, layerKeyb, layerSpace string
	var shifted bool

	//////////////
	// Keyboard //
	//////////////

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

	///////////
	// Space //
	///////////

	if m.selected == ' ' {
		if m.selected == m.requested {
			layerSpace = styleBorderCorrect.MarginLeft(19).Render(generator.Spaces(21))
		} else {
			layerSpace = styleBorderWrong.MarginLeft(19).Render(generator.Spaces(21))
		}
	} else {
		layerSpace += styleBorderNormal.MarginLeft(19).Render(generator.Spaces(21))
	}

	//////////////
	// Sentence //
	//////////////

	// Calculating the keyb Height and Width
	KeybWidth := lipgloss.Width(layerKeyb)

	// Reducing or adding to the sentence to fit the box
	if utf8.RuneCountInString(m.sentence) > KeybWidth-4 {
		layerSentence = string([]rune(m.sentence)[:KeybWidth-4])
	} else {
		layerSentence = m.sentence
		layerSentence += generator.Spaces(-utf8.RuneCountInString(m.sentence))
	}

	// Highlighting the first letter
	layerSentence = styleRequested.Render(
		string([]rune(layerSentence)[:1])) +
		string([]rune(layerSentence)[1:])

	// Adding borders
	layerSentence = styleBorderNormal.Render(layerSentence)

	//////////////////////////
	// Merging & Displaying //
	//////////////////////////

	visual := lipgloss.JoinVertical(lipgloss.Left, layerSentence, layerKeyb, layerSpace)

	// Getting visual height
	visualHeight := lipgloss.Height(visual)

	// Check if there is enough space
	if m.termWidth < KeybWidth || m.termHeight < visualHeight {
		visual = "Terminal size too small:\n"
		visual += fmt.Sprintf("Width = %d Height = %d\n\n", m.termWidth, m.termHeight)
		visual += "Needed for current config:\n"
		visual += fmt.Sprintf("Width = %d Height = %d", KeybWidth, visualHeight)
	}

	visual = lipgloss.Place(m.termWidth, m.termHeight, lipgloss.Center, lipgloss.Center, visual)
	return visual
}
