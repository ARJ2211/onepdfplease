package messages

import "github.com/chetanjangir0/onepdfplease/internal/tui/types"

type PDFOperationStatus struct {
    TaskType    string 
    Err       error
}

type Navigate struct {
	Page types.Page
}

type OutputButtonClicked struct {}
