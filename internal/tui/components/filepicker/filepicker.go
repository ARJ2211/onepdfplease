package filepicker

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chetanjangir0/onepdfplease/internal/tui/context"
	"github.com/chetanjangir0/onepdfplease/internal/tui/messages"
	"github.com/chetanjangir0/onepdfplease/internal/tui/style"
)

type Model struct {
	filepicker    filepicker.Model
	SelectedFiles []string
	ctx           *context.ProgramContext
	height        int
}

func NewModel(ctx *context.ProgramContext) Model {
	height := 20

	fp := filepicker.New()
	fp.AllowedTypes = []string{".pdf"}
	fp.CurrentDirectory, _ = os.Getwd()
	fp.SetHeight(height)
	fp.ShowPermissions = false
	// fp.KeyMap.Select = key.NewBinding(
	// 	key.WithKeys(" "),
	// 	key.WithHelp("space", "select"),
	// )
	return Model{
		filepicker: fp,
		ctx:        ctx,
		height:     height,
	}
}

func (m *Model) ClearSelected() {
	m.SelectedFiles = nil
}

func (m Model) Init() tea.Cmd {
	return m.filepicker.Init()
}

func (m *Model) SetAllowedTypes(types []string) {
	m.filepicker.AllowedTypes = types
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// case "ctrl+y":
		}
	}

	var cmd tea.Cmd
	m.filepicker, cmd = m.filepicker.Update(msg)

	// Did the user select a file?
	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
		m.SelectedFiles = append(m.SelectedFiles, path)
	}

	// Did the user select a disabled file?
	if didSelect, _ := m.filepicker.DidSelectDisabledFile(msg); didSelect {
		err := errors.New("File is not valid.")

		return m, func() tea.Msg {
			return messages.ShowError{
				Err: err,
			}
		}

	}

	return m, cmd
}

func (m Model) View() string {
	return style.RenderTwoFullCols(
		m.ctx.TermWidth,
		m.ctx.MainContentHeight,
		style.DefaultStyle.FocusedBorder,
		m.browseView(),
		m.selectedView(),
	)
}

func (m Model) browseView() string {
	var view strings.Builder
	view.WriteString("\n  ")
	view.WriteString("Pick files:")
	view.WriteString("\n\n" + m.filepicker.View() + "\n")

	return view.String()
}

func (m Model) selectedView() string {
	var view strings.Builder
	view.WriteString("\n  ")
	view.WriteString("Selected files: \n")
	view.WriteString("\n")
	for i, f := range m.SelectedFiles {
		// only show the last m.height files when files are too many
		if m.height <= len(m.SelectedFiles) && i < len(m.SelectedFiles)-m.height {
			continue
		}
		view.WriteString(m.filepicker.Styles.Selected.PaddingLeft(2).Render(filepath.Base(f)) + "\n")
	}
	view.WriteString("\n")
	return view.String()
}
