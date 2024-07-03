package main

import (
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

	// Fixed size
	layerSentence = generator.FixedSize(m.sentence, KeybWidth-4)

	// Highlighting the first letter
	layerSentence = styleRequested.Render(
		string([]rune(layerSentence)[:1])) +
		string([]rune(layerSentence)[1:])

	// Request Enter click if done
	if m.done {
		layerSentence = lipgloss.PlaceHorizontal(KeybWidth-4, lipgloss.Center,
			"Press ENTER",
		)
	}

	// Adding borders
	layerSentence = styleBorderNormal.Render(layerSentence)

	////////////////////
	// Merging layers //
	////////////////////

	visual := lipgloss.JoinVertical(lipgloss.Left, layerSentence, layerKeyb, layerSpace)

	return visual
}
