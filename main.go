package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list            list.Model
	showsTitleToURL map[string]string
	showURLToEpList map[string]string
	EpListToEpURL   map[string]string
	wantedpage      int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.wantedpage == 0 {
		items := []list.Item{}

		for k := range m.showsTitleToURL {
			newItem := item{
				title: k,
			}
			items = append(items, newItem)

		}
		// m.showsTitleToURL = mymap
		m.list.SetItems(items)
		m.wantedpage = 1
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if selected, ok := m.list.SelectedItem().(item); ok {
				// Store the selection and quit
				if m.wantedpage == 1 {
					m.showURLToEpList = getseasandeplist(m.showsTitleToURL[selected.title])
					items := []list.Item{}

					for k := range m.showURLToEpList {
						newItem := item{
							title: k,
						}
						items = append(items, newItem)

					}

					m.list.SetItems(items)
					m.wantedpage = 2

				} else if m.wantedpage == 2 {
					// fmt.Println("mpv", getURL(m.showURLToEpList[selected.title]+"2"))
					eplink := getURL(m.showURLToEpList[selected.title] + "2")
					cmd := exec.Command("mpv", "--no-terminal", eplink)
					var stderr bytes.Buffer
					cmd.Env = append(os.Environ(), "GDK_BACKEND=x11")
					cmd.Stderr = &stderr
					cmd.Env = append(os.Environ(), "DISPLAY=:0")
					if err := cmd.Start(); err != nil {
						log.Printf("MPV failed: %v | Stderr: %s", err, stderr.String())
					}
					m.wantedpage = 0
				}
			}
			m.list.ResetFilter()

		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	mymap := getlist()
	items := []list.Item{}

	for k := range mymap {
		newItem := item{
			title: k,
		}
		items = append(items, newItem)

	}
	// m.showsTitleToURL = mymap
	m := model{wantedpage: 0, showsTitleToURL: mymap, list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Shows"
	m.list.SetItems(items)
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
