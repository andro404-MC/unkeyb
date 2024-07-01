package generator

import (
	"embed"
	"math/rand"
	"strings"
)

//go:embed english-words/words.txt
var f embed.FS
var lines []string

func Load() {
	data, _ := f.ReadFile("english-words/words.txt")
	lines = strings.Split(string(data), "\n")
}

func Sentence() string {
	var s string
	wrdCnt := rand.Intn(30) + 21

	for i := 0; i < wrdCnt; i++ {
		s += lines[rand.Intn(len(lines))]
		if i+1 != wrdCnt {
			s += " "
		}
	}

	return s
}
