package main

import "github.com/charmbracelet/lipgloss"

type model struct {
	requested rune
	selected  rune
	sentence  string

	layout  string
	shifted bool

	termWidh   int
	termHeight int
}

type row struct {
	prefix  string
	postfix string

	sKeys []rune
	keys  []rune
}

var keyList []rune

const (
	colorCorrect   = "\033[38;5;8m"
	colorWrong     = "\033[38;5;1m"
	colorRequested = "\033[4m"
	colorReset     = "\033[0m"
)

var (
	styleBorderNormal = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				PaddingLeft(1).PaddingRight(1)

	styleBorderCorrect = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				PaddingLeft(1).PaddingRight(1).
				Foreground(lipgloss.ANSIColor(8))

	styleBorderWrong = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				PaddingLeft(1).PaddingRight(1).
				Foreground(lipgloss.ANSIColor(1))
)

func generateList(layout string) {
	for _, v := range layouts[layout] {
		keyList = append(keyList, v.keys...)
		keyList = append(keyList, v.sKeys...)
	}
}
