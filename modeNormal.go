package main

import (
	"fmt"
	"unicode/utf8"
)

func normalKeyb(m *model) string {
	var s string
	var sentence string
	if utf8.RuneCountInString(m.sentence) > 39 {
		sentence += string([]rune(m.sentence)[:39])
	} else {
		sentence += m.sentence
	}

	s += fmt.Sprintf(
		"\n     %s\n\n",
		colorRequested+string([]rune(sentence)[:1])+colorReset+string([]rune(sentence)[1:]),
	)

	for _, item := range layouts[m.layout] {
		for _, shiftedKey := range item.sKeys {
			if shiftedKey == m.selected {
				m.shifted = true
				break
			}
		}
	}

	for _, v := range layouts[m.layout] {
		// prefix
		s += v.prefix

		var rangedSlice *[]rune

		if m.shifted {
			rangedSlice = &v.sKeys
		} else {
			rangedSlice = &v.keys
		}

		// keys
		for _, k := range *rangedSlice {
			isClicked := m.selected == k

			if isClicked {
				if k == m.requested {
					s += fmt.Sprintf("%s%c%s  ", colorCorrect, k, colorReset)
				} else {
					s += fmt.Sprintf("%s%c%s  ", colorWrong, k, colorReset)
				}
			} else {
				s += fmt.Sprintf("%c  ", k)
			}
		}

		// postfix
		s += v.postfix
	}

	// space
	if m.selected == ' ' {
		if m.selected == m.requested {
			s += fmt.Sprintf("\n\t    %s🬞🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬏%s", colorCorrect, colorReset)
			s += fmt.Sprintf("\n\t    %s🬁🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬀%s", colorCorrect, colorReset)
		} else {
			s += fmt.Sprintf("\n\t    %s🬞🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬏%s", colorWrong, colorReset)
			s += fmt.Sprintf("\n\t    %s🬁🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬀%s", colorWrong, colorReset)
		}
	} else {
		s += "\n\t    🬞🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬭🬏"
		s += "\n\t    🬁🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬂🬀"
	}

	return s
}
