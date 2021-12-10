package usecase

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/file_manager"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/spf13/viper"
)

type FileUsecase struct {
	repo file_manager.IFileRepo
}

func NewFileUsecase(repo file_manager.IFileRepo) *FileUsecase {
	return &FileUsecase{
		repo: repo,
	}
}

func (u *FileUsecase) CreateFiles(files []*internal_models.File, ticketID uint64) models.Err {
	defaultPath := viper.GetString("app.store.path")
	year, month, day := time.Now().Date()
	pathToFolder := buildPathToFolder(defaultPath, year, int(month), day)

	for _, file := range files {
		file.TicketId = ticketID
		file.Date = time.Now().Truncate(time.Second)
		file.Extension = path.Ext(file.Name)
		file.Name = strings.TrimSuffix(file.Name, file.Extension)
		file.Path = fmt.Sprintf("%s/%s%s", pathToFolder, file.Name, file.Extension)

		err := u.repo.CreateFile(file)
		if err != nil {
			return models.InternalError(err.Error())
		}
	}

	return nil
}

func buildPathToFolder(defaultPath string, year, month, day int) string {
	pathToDay := fmt.Sprintf("%s/%d/%d/%d", defaultPath, year, month, day)
	if _, err := os.Stat(pathToDay); os.IsNotExist(err) {
		pathToMonth := fmt.Sprintf("%s/%d/%d", defaultPath, year, month)
		if _, err := os.Stat(pathToMonth); os.IsExist(err) {
			os.Mkdir(pathToDay, 0777)
		} else {
			pathToYear := fmt.Sprintf("%s/%d", defaultPath, year)
			if _, err := os.Stat(pathToYear); os.IsExist(err) {
				os.Mkdir(pathToMonth, 0777)
				os.Mkdir(pathToDay, 0777)
			} else {
				os.Mkdir(pathToYear, 0777)
				os.Mkdir(pathToMonth, 0777)
				os.Mkdir(pathToDay, 0777)
			}
		}
	}
	return pathToDay
}

func (u *FileUsecase) GetFile(fileID uint64) (*internal_models.File, models.Err) {
	file, err := u.repo.GetFile(fileID)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	return file, nil
}

func (u *FileUsecase) GetTicketFiles(ticketID uint64) ([]*internal_models.File, models.Err) {
	files, err := u.repo.GetTicketFiles(ticketID)
	if err != nil {
		return nil, models.InternalError(err.Error())
	}

	return files, nil
}
