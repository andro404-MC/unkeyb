package main

var layouts = map[string][]row{
	"qwerty-uk": {
		{
			sKeys: []rune{'¬', '!', '"', '£', '$', '%', '^', '&', '*', '(', ')', '_', '+'},
			keys:  []rune{'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
		},
		{
			sKeys: []rune{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '{', '}'},
			keys:  []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']'},
		},
		{
			sKeys: []rune{'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ':', '@', '~'},
			keys:  []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\'', '#'},
		},
		{
			sKeys: []rune{'|', 'Z', 'X', 'C', 'V', 'B', 'N', 'M', '<', '>', '?'},
			keys:  []rune{'\\', 'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/'},
		},
	},

	"qwerty": {
		{
			sKeys: []rune{'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+'},
			keys:  []rune{'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
		},
		{
			sKeys: []rune{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '{', '}', '|'},
			keys:  []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']', '\\'},
		},
		{
			sKeys: []rune{'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ':', '"'},
			keys:  []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\''},
		},
		{
			sKeys: []rune{'Z', 'X', 'C', 'V', 'B', 'N', 'M', '<', '>', '?'},
			keys:  []rune{'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/'},
		},
	},

	"dvorak": {
		{
			sKeys: []rune{'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '{', '}'},
			keys:  []rune{'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '[', ']'},
		},
		{
			sKeys: []rune{'"', '<', '>', 'P', 'Y', 'F', 'G', 'C', 'R', 'L', '?', '+', '|'},
			keys:  []rune{'\'', ',', '.', 'p', 'y', 'f', 'g', 'c', 'r', 'l', '/', '=', '\\'},
		},
		{
			sKeys: []rune{'A', 'O', 'E', 'U', 'I', 'D', 'H', 'T', 'N', 'S', '_'},
			keys:  []rune{'a', 'o', 'e', 'u', 'i', 'd', 'h', 't', 'n', 's', '-'},
		},
		{
			sKeys: []rune{':', 'Q', 'J', 'K', 'X', 'B', 'M', 'W', 'V', 'Z'},
			keys:  []rune{';', 'q', 'j', 'k', 'x', 'b', 'm', 'w', 'v', 'z'},
		},
	},

	"colemak": {
		{
			sKeys: []rune{'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+'},
			keys:  []rune{'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
		},
		{
			sKeys: []rune{'Q', 'W', 'F', 'P', 'G', 'J', 'L', 'U', 'Y', ':', '{', '}', '|'},
			keys:  []rune{'q', 'w', 'f', 'p', 'g', 'j', 'l', 'u', 'y', ';', '[', ']', '\\'},
		},
		{
			sKeys: []rune{'A', 'R', 'S', 'T', 'D', 'H', 'N', 'E', 'I', 'O', '"'},
			keys:  []rune{'a', 'r', 's', 't', 'd', 'h', 'n', 'e', 'i', 'o', '\''},
		},
		{
			sKeys: []rune{'Z', 'X', 'C', 'V', 'B', 'K', 'M', '<', '>', '?'},
			keys:  []rune{'z', 'x', 'c', 'v', 'b', 'k', 'm', ',', '.', '/'},
		},
	},

	"colemak_dh": {
		{
			sKeys: []rune{'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+'},
			keys:  []rune{'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '='},
		},
		{
			sKeys: []rune{'Q', 'W', 'F', 'P', 'B', 'J', 'L', 'U', 'Y', ':', '{', '}', '|'},
			keys:  []rune{'q', 'w', 'f', 'p', 'b', 'j', 'l', 'u', 'y', ';', '[', ']', '\\'},
		},
		{
			sKeys: []rune{'A', 'R', 'S', 'T', 'G', 'M', 'N', 'E', 'I', 'O', '"'},
			keys:  []rune{'a', 'r', 's', 't', 'g', 'm', 'n', 'e', 'i', 'o', '\''},
		},
		{
			sKeys: []rune{'X', 'C', 'D', 'V', 'Z', 'K', 'H', '<', '>', '?'},
			keys:  []rune{'x', 'c', 'd', 'v', 'z', 'k', 'h', ',', '.', '/'},
		},
	},

	"azerty": {
		{
			sKeys: []rune{'~', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '°', '+'},
			keys:  []rune{'`', '&', 'é', '"', '\'', '(', '-', 'è', '_', 'ç', 'à', ')', '='},
		},
		{
			sKeys: []rune{'A', 'Z', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '¨', '£'},
			keys:  []rune{'a', 'z', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '^', '$'},
		},
		{
			sKeys: []rune{'Q', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', '%', 'µ'},
			keys:  []rune{'q', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'ù', '*'},
		},
		{
			sKeys: []rune{'>', 'W', 'X', 'C', 'V', 'B', 'N', '?', '.', '/', '§'},
			keys:  []rune{'<', 'w', 'x', 'c', 'v', 'b', 'n', ',', ';', ':', '!'},
		},
	},
}
