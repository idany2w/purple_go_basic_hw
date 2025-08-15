package file

import (
	"os"
	"regexp"
)

func ReadFIle(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func IsJsonFileName(name string) bool {
	isJson, _ := regexp.MatchString(`\.json$`, name)
	return isJson
}
