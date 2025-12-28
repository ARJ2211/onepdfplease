package types

type Page int

const (
	MenuPage Page = iota
	MergePage
	SplitPage
	EncryptPage
)

type NavigateMsg struct {
	Page Page
}

type QuitFilePickerMsg struct {
	Paths []string
}
