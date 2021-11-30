package dto

import (
	"fmt"
	"time"

	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
)

type OutFile struct {
	ID   uint64    `json:"file_id"`
	Name string    `json:"file_name"`
	Data string    `json:"file_data"`
	Date time.Time `json:"file_date"`
}

func ToOutFile(file *internal_models.File) OutFile {
	return OutFile{
		ID:   file.ID,
		Name: fmt.Sprint(file.Name + file.Extension),
		Data: file.Data,
		Date: file.Date,
	}
}

func ToModelFile(file OutFile) *internal_models.File {
	return &internal_models.File{
		Name: file.Name,
		Data: file.Data,
	}
}
