package main

import (
	"os"
	"path/filepath"

	// "strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	items    []string
	selected int
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) View() string {
	selectedStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	var output string = ""
	for index, item := range m.items {
		if index == m.selected {

			output += selectedStyle.Render(">"+item) + "\n"
		} else {
			output += item + "\n"

		}
	}
	return output
}
func (m model) Update(msgType tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msgType.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.selected > 0 {
				m.selected--
			}

		case "down":
			if m.selected < len(m.items)-1 {
				m.selected++
			}

		case "q":
			return m, tea.Quit
		}

	}
	return m, nil
}
func isParent(index, length int) string {

	if index == length {
		return "└──"
	} else {
		return "├──"
	}
}

func parseTree(path string, depth int) {
	data, err := os.ReadDir(path)
	if err != nil {
		println(err.Error())
		return
	}

	for index, entry := range data {

		if entry.IsDir() {

			println(strings.Repeat("\t", depth), isParent(index, len(data)-1)+entry.Name()+"/")
			parseTree(filepath.Join(path, entry.Name()), depth+1)
		} else {
			println(strings.Repeat("\t", depth), isParent(index, len(data)-1)+entry.Name())

		}
	}
}
func main() {
	// path := os.Args[1]

	// parseTree(path, 0)
	p := tea.NewProgram(model{items: []string{"src", "bb", "app", "go.sum"}, selected: 0})
	p.Run()
}
