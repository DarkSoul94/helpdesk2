package mysql

import "github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"

func (r *FileRepo) toModelFile(file dbFile) *internal_models.File {
	mFile := &internal_models.File{
		ID:        file.ID,
		Name:      file.Name,
		Date:      file.Date,
		Extension: file.Extension,
		TicketId:  file.TicketId,
	}

	if file.Data.Valid {
		mFile.Data = file.Data.String
	}

	if file.Path.Valid {
		mFile.Path = file.Path.String
	}

	return mFile
}

func (r *FileRepo) toDbFile(file *internal_models.File) dbFile {
	dbFile := dbFile{
		ID:        file.ID,
		Name:      file.Name,
		Date:      file.Date,
		Extension: file.Extension,
		TicketId:  file.TicketId,
	}

	if len(file.Data) > 0 {
		dbFile.Data.String = file.Data
		dbFile.Data.Valid = true
	} else {
		dbFile.Data.Valid = false
	}

	if len(file.Path) > 0 {
		dbFile.Path.String = file.Path
		dbFile.Path.Valid = true
	} else {
		dbFile.Path.Valid = false
	}

	return dbFile
}
