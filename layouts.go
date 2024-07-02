package main

var layouts = map[string][]row{
	"gb": {
		{
			prefix: 0,
			sKeys:  []rune{'¬', '!', '"', '£', '$', '%', '^', '&', '*', '(', ')', '_', '+'},
			keys:   []rune{'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
		},
		{
			prefix: 4,
			sKeys:  []rune{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '{', '}'},
			keys:   []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']'},
		},
		{
			prefix: 5,
			sKeys:  []rune{'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ':', '@', '~'},
			keys:   []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\'', '#'},
		},
		{
			prefix: 4,
			sKeys:  []rune{'|', 'Z', 'X', 'C', 'V', 'B', 'N', 'M', '<', '>', '?'},
			keys:   []rune{'\\', 'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/'},
		},
	},
}
