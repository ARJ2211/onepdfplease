package merge

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	add    key.Binding
	remove key.Binding
	merge  key.Binding
	save   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.add, k.remove, k.merge, k.save}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.add, k.remove, k.merge, k.save}, // first column
		// {k.Help, k.Quit},                // second column
	}
}

var keys = keyMap{
	add: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "Add files"),
	),
	remove: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "Remove files"),
	),
	merge: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "Merge PDFs"),
	),
	save: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "Save PDF"),
	),
}
