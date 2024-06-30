package main

type model struct {
	requested rune
	next      rune
	selected  rune

	layout  string
	shifted bool
}

var keyList []rune

type row struct {
	prefix  string
	postfix string

	sKeys []rune
	keys  []rune
}

func generateList(layout string) {
	for _, v := range layouts[layout] {
		keyList = append(keyList, v.keys...)
		keyList = append(keyList, v.sKeys...)
	}
}
