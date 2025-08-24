package files

import (
	"errors"
	"os"
	"path/filepath"
)

const jsonExtension = ".json"

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	file, err := os.ReadFile(db.filename)

	if err != nil {
		return nil, errors.New("При чтении файла произошла ошибка")
	}

	if filepath.Ext(db.filename) != jsonExtension {
		return nil, errors.New("Невалидный json-файл")
	}

	return file, nil
}

func (db *JsonDb) Write(content []byte) (bool, error) {
	file, err := os.Create(db.filename)

	if err != nil {
		return false, err
	}

	_, err = file.Write(content)

	defer file.Close()

	if err != nil {
		return false, err
	}

	return true, nil
}
