package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"unkeyb/generator"
)

func uiMinimal(m *model) string {
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
					keys = append(keys, styleCorrect.Render(string(k)))
				} else {
					keys = append(keys, styleWrong.Render(string(k)))
				}
			} else {
				keys = append(keys, styleNormal.Render(string(k)))
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

	spaceShape := "┌──────────────┐\n"
	spaceShape += "└──────────────┘"

	if m.selected == ' ' {
		if m.selected == m.requested {
			layerSpace = styleCorrect.Render(spaceShape)
		} else {
			layerSpace = styleWrong.Render(spaceShape)
		}
	} else {
		layerSpace += styleNormal.Render(spaceShape)
	}

	layerSpace = lipgloss.PlaceHorizontal(KeybWidth, lipgloss.Center, layerSpace)

	//////////////
	// Sentence //
	//////////////

	// Fixed size
	layerSentence = generator.FixedSize(m.sentence, KeybWidth)

	// Highlighting the first letter
	layerSentence = styleRequested.Render(
		string([]rune(layerSentence)[:1])) +
		string([]rune(layerSentence)[1:])

	// Request Enter click if done
	if m.done {
		layerSentence = lipgloss.PlaceHorizontal(KeybWidth, lipgloss.Center,
			fmt.Sprintf("WPM:%.2f Press Enter", m.wpm),
		)
	}

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
