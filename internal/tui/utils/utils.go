package utils

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chetanjangir0/onepdfplease/internal/tui/messages"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func Merge(inFiles []string, outFile string) tea.Cmd {
	return func() tea.Msg {
		tasktype := "Merge"
		if len(inFiles) == 0 {
			return messages.PDFOperationStatus{
				TaskType: tasktype,
				Err:      fmt.Errorf("There are no files to merge"),
			}
		}
		err := api.MergeCreateFile(inFiles, outFile, false, nil)
		if err != nil {
			return messages.PDFOperationStatus{
				TaskType: tasktype,
				Err:      err,
			}
		}

		return messages.PDFOperationStatus{
			TaskType: tasktype,
		}
	}
}

func Encrypt(inFile, outFile string, password string) tea.Cmd {
	return func() tea.Msg {
		taskType := "Encrypt"
		if len(password) == 0 {
			return messages.PDFOperationStatus{
				TaskType: taskType,
				Err:      fmt.Errorf("Please provide a password"),
			}

		}
		conf := model.NewDefaultConfiguration()
		conf.UserPW = password
		conf.OwnerPW = password

		err := api.EncryptFile(inFile, outFile, conf)
		if err != nil {
			return messages.PDFOperationStatus{
				TaskType: taskType,
				Err:      err,
			}
		}
		return messages.PDFOperationStatus{
			TaskType: taskType,
		}

	}

}
