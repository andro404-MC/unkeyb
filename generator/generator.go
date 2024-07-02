package generator

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed google-10000-english/google-10000-english-no-swears.txt
var f embed.FS
var lines []string

const AnsiReset = "\033[0m"

func Load() {
	data, _ := f.ReadFile("google-10000-english/google-10000-english-no-swears.txt")
	lines = strings.Split(string(data), "\n")
}

func Sentence() string {
	var s string
	wrdCnt := rand.Intn(30) + 21

	for i := 0; i < wrdCnt; i++ {
		s += lines[rand.Intn(len(lines))]
		if i+1 != wrdCnt {
			if cos := rand.Intn(5); cos == 0 {
				s += ", "
			} else {
				s += " "
			}
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
