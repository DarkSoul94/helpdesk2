package mysql

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/DarkSoul94/helpdesk2/pkg/logger"
	"github.com/DarkSoul94/helpdesk2/pkg_ticket/internal_models"
	"github.com/jmoiron/sqlx"
)

func NewFileRepo(db *sql.DB) *FileRepo {
	return &FileRepo{
		db: sqlx.NewDb(db, "mysql"),
	}
}

func (r *FileRepo) CreateFile(file *internal_models.File) error {
	var (
		id     uint64
		osFile *os.File
		data   []byte
		query  string
		err    error
	)

	query = `SELECT file_id FROM files
				WHERE file_name = ?
				AND file_extension = ?
				AND ticket_id = ?
				LIMIT 1`

	err = r.db.Get(&id, query, file.Name, file.Extension, file.TicketId)
	if err != nil {
		query = `INSERT INTO files SET
					ticket_id = :ticket_id,
					file_name = :file_name,
					file_extension = :file_extension,
					file_date = :file_date,
					file_data = :file_data,
					path = :path`

		osFile, err = os.Create(file.Path)
		if err != nil {
			logger.LogError(
				"Failed create file in directory",
				"pkg_ticket/file_manager/repo/mysql",
				fmt.Sprintf("path: %s", file.Path),
				err,
			)
			return err
		}
	} else {
		query = `UPDATE files SET
					file_date = :file_date,
					file_data = :file_data,
					path = :path
					WHERE file_id = :file_id`

		osFile, err = os.Open(file.Path)
		if err != nil {
			logger.LogError(
				"Failed open file from directory",
				"pkg_ticket/file_manager/repo/mysql",
				fmt.Sprintf("path: %s", file.Path),
				err,
			)
			return err
		}
	}

	defer osFile.Close()

	split := strings.Split(file.Data, ",")
	data, err = base64.StdEncoding.DecodeString(split[1])
	if err != nil {
		logger.LogError(
			"Failed decode file",
			"pkg_ticket/file_manager/repo/mysql",
			fmt.Sprintf("file name: %s", file.Name),
			err,
		)
		return err
	}

	osFile.Write(data)
	file.Data = split[0]

	_, err = r.db.NamedExec(query, r.toDbFile(file))
	if err != nil {
		logger.LogError(
			"Failed create file",
			"pkg_ticket/file_manager/repo/mysql",
			fmt.Sprintf("file name: %s; file extension: %s; file date: %s; path: %s; ticket id: %d;", file.Name, file.Extension, file.Date, file.Path, file.TicketId),
			err,
		)
		return err
	}

	return nil
}

func (r *FileRepo) GetFile(fileID uint64) (*internal_models.File, error) {
	var (
		dbFile   dbFile
		file     *os.File
		fileData []byte
		query    string
		err      error
	)

	query = `SELECT * FROM files WHERE file_id = ?`
	err = r.db.Get(&dbFile, query, fileID)
	if err != nil {
		logger.LogError(
			"Failed read file from db",
			"pkg_ticket/file_manager/repo/mysql",
			fmt.Sprintf("file id: %d;", fileID),
			err,
		)
		return nil, err
	}

	if dbFile.Path.Valid {
		file, err = os.Open(dbFile.Path.String)
		if err != nil {
			logger.LogError(
				"Failed open file in directory",
				"pkg_ticket/file_manager/repo/mysql",
				fmt.Sprintf("file path: %s; file id: %d;", dbFile.Path.String, fileID),
				errors.New("Файл удален с сервера"),
			)
			return nil, err
		}
		defer file.Close()

		packet := make([]byte, 10000)

		for {
			byteCount, err := file.Read(packet)
			fileData = append(fileData, packet[:byteCount]...)
			if err == io.EOF { // если конец файла
				break // выходим из цикла
			}
		}

		encodeStr := base64.StdEncoding.EncodeToString(fileData)

		dbFile.Data.String += "," + encodeStr
		dbFile.Data.Valid = true
	}

	return r.toModelFile(dbFile), nil
}

func (r *FileRepo) GetTicketFiles(ticketID uint64) ([]*internal_models.File, error) {
	var (
		dbFiles []dbFile
		mFiles  []*internal_models.File
		query   string
		err     error
	)

	query = `SELECT file_id, file_name, file_date FROM files WHERE ticket_id = ?`
	err = r.db.Select(&dbFiles, query, ticketID)
	if err != nil {
		logger.LogError(
			"Failed read ticket files",
			"pkg_ticket/file_manager/repo/mysql",
			fmt.Sprintf("ticket id: %d;", ticketID),
			err,
		)
		return nil, err
	}

	for _, file := range dbFiles {
		mFiles = append(mFiles, r.toModelFile(file))
	}

	return mFiles, nil
}

func (r *FileRepo) Close() error {
	return r.db.Close()
}
