package context

import "github.com/chetanjangir0/onepdfplease/internal/tui/types"

type ProgramContext struct {
	MainContentHeight int
	TermWidth         int
	TermHeight        int
	Status            string
	StatusType        StatusType
	CurrentPage       types.Page
	// Config            *config.Config
}

type StatusType int

const (
	Error StatusType = iota
	Success
	Processing
	None
)

func (m *ProgramContext) SetStatusProcessing(msg string) {
	m.Status = msg
	m.StatusType = Processing
}

func (m *ProgramContext) SetStatusError(msg string) {
	m.Status = msg
	m.StatusType = Error
}

func (m *ProgramContext) SetStatusSuccess(msg string) {
	m.Status = msg
	m.StatusType = Success
}

func (m *ProgramContext) ClearStatus() {
	m.Status = ""
	m.StatusType = None
}
