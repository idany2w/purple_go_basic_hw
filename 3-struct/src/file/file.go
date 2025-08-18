package file

import (
	"errors"
	"os"
	"regexp"
)

type FileProvider struct {
	storagePath string
}

func NewFileProvider(storagePath string) (*FileProvider, error) {
	if !IsJsonFileName(storagePath) {
		return nil, errors.New("INVALID_FILE_NAME")
	}

	return &FileProvider{storagePath: storagePath}, nil
}

func (f *FileProvider) Read() ([]byte, error) {
	data, err := os.ReadFile(f.storagePath)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func IsJsonFileName(name string) bool {
	isJson, _ := regexp.MatchString(`\.json$`, name)
	return isJson
}

func (f *FileProvider) Save(data []byte) error {
	return os.WriteFile(f.storagePath, data, 0755)
}
