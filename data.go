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
	keys    []rune
	sKeys   []rune
}

var layouts = map[string][]row{
	"gb": {
		{
			prefix: "\n\t", postfix: "\n\n",
			keys:  []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
			sKeys: []rune{'!', '"', '£', '$', '%', '^', '&', '*', '(', ')', '_', '+'},
		},
		{
			prefix: "\t ", postfix: "\n\n",
			keys:  []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']'},
			sKeys: []rune{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '{', '}'},
		},
		{
			prefix: "\t  ", postfix: "\n\n",
			keys:  []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\'', '#'},
			sKeys: []rune{'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ':', '@', '~'},
		},
		{
			prefix: "\t ", postfix: "\n",
			keys:  []rune{'\\', 'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/'},
			sKeys: []rune{'|', 'Z', 'X', 'C', 'V', 'B', 'N', 'M', '<', '>', '?'},
		},
	},
}
