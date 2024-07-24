package main

import (
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	layout   string
	sentence string
	minimal  bool

	runeCount int
	startTime int64
	wpm       float32

	requested rune
	selected  rune

	termWidth  int
	termHeight int

	done     bool
	fistChar bool
	prevKey string
}

type row struct {
	sKeys []rune
	keys  []rune
}

var keyList []rune

var (
	styleNormal = lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).PaddingTop(1)

	styleCorrect = lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).PaddingTop(1).
			Foreground(lipgloss.ANSIColor(4))

	styleWrong = lipgloss.NewStyle().
			PaddingLeft(1).PaddingRight(1).PaddingTop(1).
			Foreground(lipgloss.ANSIColor(1))

	styleBorderNormal = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				PaddingLeft(1).PaddingRight(1)

	styleBorderCorrect = lipgloss.NewStyle().
				BorderStyle(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.ANSIColor(4)).
				PaddingLeft(1).PaddingRight(1).
				Foreground(lipgloss.ANSIColor(4))

	styleBorderWrong = lipgloss.NewStyle().
				BorderStyle(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.ANSIColor(1)).
				PaddingLeft(1).PaddingRight(1).
				Foreground(lipgloss.ANSIColor(1))

	styleRequested = lipgloss.NewStyle().
			Underline(true)
)

func generateList(layout string) {
	for _, v := range layouts[layout] {
		keyList = append(keyList, v.keys...)
		keyList = append(keyList, v.sKeys...)
	}
}
