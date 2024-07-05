package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"unkeyb/generator"
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

	var KeybWidth int

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

		row := lipgloss.JoinHorizontal(lipgloss.Right, keys...)

		// Merging to row
		rows = append(rows,
			lipgloss.JoinHorizontal(lipgloss.Right, keys...),
		)

		rowWidth := lipgloss.Width(row)

		if rowWidth > KeybWidth {
			KeybWidth = rowWidth
		}
	}

	for i := 0; i < len(rows); i++ {
		rows[i] = lipgloss.PlaceHorizontal(KeybWidth, lipgloss.Center, rows[i])
	}

	// Mergin rows to the Keyboard layer
	layerKeyb = lipgloss.JoinVertical(lipgloss.Left, rows...)

	///////////
	// Space //
	///////////

	if m.selected == ' ' {
		if m.selected == m.requested {
			layerSpace = styleBorderCorrect.Render(generator.Spaces(21))
		} else {
			layerSpace = styleBorderWrong.Render(generator.Spaces(21))
		}
	} else {
		layerSpace += styleBorderNormal.Render(generator.Spaces(21))
	}

	layerSpace = lipgloss.PlaceHorizontal(KeybWidth, lipgloss.Center, layerSpace)

	//////////////
	// Sentence //
	//////////////

	// Fixed size
	layerSentence = generator.FixedSize(m.sentence, KeybWidth-4)

	// Highlighting the first letter
	layerSentence = styleRequested.Render(
		string([]rune(layerSentence)[:1])) +
		string([]rune(layerSentence)[1:])

	// Request Enter click if done
	if m.done {
		layerSentence = lipgloss.PlaceHorizontal(KeybWidth-4, lipgloss.Center,
			fmt.Sprintf("WPM:%.2f Press Enter", m.wpm),
		)
	}

	// Adding borders
	layerSentence = styleBorderNormal.Render(layerSentence)

	////////////////////
	// Merging layers //
	////////////////////

	visual := lipgloss.JoinVertical(
		lipgloss.Left,
		layerSentence,
		layerKeyb,
		layerSpace,
	)

	return visual
}
