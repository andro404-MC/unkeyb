package main

type model struct {
	requested rune
	selected  rune
	sentence  string

	layout  string
	shifted bool
}

type row struct {
	prefix  string
	postfix string

	sKeys []rune
	keys  []rune
}

var keyList []rune

const (
	colorCorrect = "\033[38;5;49m"
	colorWrong   = "\033[38;5;196m"
	colorReset   = "\033[0m"
)

func generateList(layout string) {
	for _, v := range layouts[layout] {
		keyList = append(keyList, v.keys...)
		keyList = append(keyList, v.sKeys...)
	}
}
