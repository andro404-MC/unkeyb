package generator

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"
)

//go:embed wordlists/*
var f embed.FS

var (
	short  []string
	medium []string
	long   []string
)

const AnsiReset = "\033[0m"

func Load(lang string) {
	data, _ := f.ReadFile("wordlists/" + lang + "/short.txt")
	short = strings.Split(string(data), "\n")

	data, _ = f.ReadFile("wordlists/" + lang + "/medium.txt")
	medium = strings.Split(string(data), "\n")

	data, _ = f.ReadFile("wordlists/" + lang + "/long.txt")
	long = strings.Split(string(data), "\n")
}

func Sentence() string {
	var s string
	wrdCnt := rand.Intn(20-10) + 10

	wasShort := true
	for i := 0; i < wrdCnt; i++ {
		if wasShort {
			r := rand.Intn(2)
			if r == 1 {
				s += medium[rand.Intn(len(medium))]
			} else {
				s += long[rand.Intn(len(long))]
			}
		} else {
			s += short[rand.Intn(len(short))]
		}
		wasShort = !wasShort

		if i+1 != wrdCnt {
			s += " "
		} else {
			s += "."
		}
	}

	return s
}

func Spaces(count int) string {
	var s string
	for i := 0; i < count; i++ {
		s += " "
	}
	return s
}

func AnsiToString(num uint) string {
	return fmt.Sprintf("\033[38;5;%dm", num)
}

func FixedSize(text string, size int) string {
	var s string
	if utf8.RuneCountInString(text) > size {
		s = string([]rune(text)[:size])
	} else {
		s = text
		s += Spaces(size - utf8.RuneCountInString(text))
	}
	return s
}
