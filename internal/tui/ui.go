package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chetanjangir0/onepdfplease/internal/tui/pages/menu"
)

type Page int

const (
	Menu Page = iota
	Merge
	Split
)

type model struct {
	quitting    bool
	currentPage Page

	// each page has its own model
	menuModel menu.Model
}

func InitialModel() model {

	return model{
		currentPage: Menu,
		menuModel:   menu.NewModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	switch m.currentPage {
	case Menu:
		m.menuModel, cmd = m.menuModel.Update(msg)
	}

	return m, cmd 
}

func (m model) View() string {
	switch m.currentPage {
	case Menu:
		return m.menuModel.View()
	}
	return ""

}
