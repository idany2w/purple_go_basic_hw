package storage

import (
	"demo/struct/src/bins"
	"encoding/json"
	"os"
)

func ReadBinList(name string) (bins.BinList, error) {
	var bl bins.BinList
	data, err := readFileData(name)

	if err != nil {
		return bl, err
	}

	err = json.Unmarshal(data, &bl)

	if err != nil {
		return bl, err
	}

	return bl, nil
}

func SaveBinList(bl bins.BinList, name string) error {
	data, err := json.Marshal(bl)

	if err != nil {
		return err
	}

	return saveDataToFile(data, name)
}

func saveDataToFile(data []byte, fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		return err
	}

	return nil
}

func readFileData(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return data, nil
}
