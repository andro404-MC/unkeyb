package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"unkeyb/generator"
)

func main() {
	var leyNames, lang string

	// Initializing the model
	m := model{}

	// Getting Layout names
	for k := range layouts {
		leyNames += k + ","
	}

	// Trimming last char
	leyNames = strings.TrimSuffix(leyNames, ",")

	// Handling flags
	flag.StringVar(&lang, "l", "en", "Language (en,fr)")
	flag.StringVar(&m.layout, "k", "qwerty", fmt.Sprintf("layout (%s)", leyNames))
	flag.BoolVar(&m.minimal, "m", false, "enable minimal UI")
	flag.Parse()

	// Load the language
	generator.Load(lang)

	// generate keys
	generateList(m.layout)

	m.fistChar = true
	m.sentence = generator.Sentence()
	m.runeCount = utf8.RuneCountInString(m.sentence)

	// Stating tea loop
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("WHAAAAAT ITS BROKEN ALREAAADY ???\ndetails: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Checking msg type
	switch msg := msg.(type) {
	// Keyboard input
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyTab:
			if !m.done {
				m.prevKey = "tab"
			}

		case tea.KeyEnter:
			// new Sentence creation and States reset
			if m.done || m.prevKey == "tab"{
				m.sentence = generator.Sentence()
				m.runeCount = utf8.RuneCountInString(m.sentence)
				m.done = false
				m.fistChar = true
				m.prevKey = ""
			}


		case tea.KeyEscape, tea.KeyCtrlC:
			// bay bay
			return m, tea.Quit

		case tea.KeyRunes, tea.KeySpace:
			if !m.done {
				// set the time when the first character is typed
				if m.fistChar {
					m.startTime = time.Now().Unix()
					m.fistChar = false
				}

				// Register the typed character
				m.selected = msg.Runes[0]

				// Set the Requested character
				m.requested = []rune(m.sentence)[0]

				// Delete the first character if correct
				if m.selected == m.requested {
					m.sentence = strings.TrimPrefix(m.sentence, string([]rune(m.sentence)[0]))
				}

				// Calcuating wpm
				if m.sentence == "" {
					bTime := float32(time.Now().Unix()-m.startTime) / 60
					m.wpm = float32((float32(m.runeCount) / 5) / bTime)
					m.done = true
				}
			}
		}

	// Terminal resize
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
	}



	return m, nil
}

func (m model) View() string {
	// Getting the template
	var visual string
	if m.minimal {
		visual = uiMinimal(&m)
	} else {
		visual = uiNormal(&m)
	}

	// Getting visual height & width
	visualWidth := lipgloss.Width(visual)
	visualHeight := lipgloss.Height(visual)

	// Check if there is enough space
	if m.termWidth < visualWidth || m.termHeight < visualHeight {
		visual = "Terminal size too small:\n"

		// Coloring
		// Width
		if m.termWidth < visualWidth {
			visual += fmt.Sprintf("Width = %s%d%s",
				generator.AnsiToString(1), m.termWidth, generator.AnsiReset)
		} else {
			visual += fmt.Sprintf("Width = %s%d%s",
				generator.AnsiToString(2), m.termWidth, generator.AnsiReset)
		}
		// Height
		if m.termHeight < visualHeight {
			visual += fmt.Sprintf(" Height = %s%d%s\n\n",
				generator.AnsiToString(1), m.termHeight, generator.AnsiReset)
		} else {
			visual += fmt.Sprintf(" Height = %s%d%s\n\n",
				generator.AnsiToString(2), m.termHeight, generator.AnsiReset)
		}

		// Required size
		visual += "Needed for current config:\n"
		visual += fmt.Sprintf("Width = %d Height = %d", visualWidth, visualHeight)
	}

	// Centering
	visual = lipgloss.Place(m.termWidth, m.termHeight, lipgloss.Center, lipgloss.Center, visual)
	return visual
}
