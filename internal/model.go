package internal

import tea "github.com/charmbracelet/bubbletea"

type MenuState int

const (
	Tools MenuState = iota
	Picker
)

type model struct {
	cursor       int
	CurrentMenu  MenuState
	Tools        []string
	SelectedTool string
	status       string
	width        int
	height       int
}

func InitialModel() model {

	return model{
		cursor:      0,
		CurrentMenu: Tools,
		Tools:       []string{"Merge PDF", "Split PDF"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
