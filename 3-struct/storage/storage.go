package storage

import (
	"demo/go-json/bins"
	"demo/go-json/files"
	"encoding/json"
	"time"
)

type Storage struct {
	UpdatedAt time.Time    `json:"updatedAt"`
	BinList   bins.BinList `json:"binList"`
}

const fileName = "bins.json"

func NewStorage() *Storage {
	db := files.NewJsonDb(fileName)

	data, err := db.Read()

	if err != nil {
		panic(err)
	}

	var storage Storage

	err = json.Unmarshal(data, &storage)

	if err != nil {
		panic(err)
	}

	return &storage
}

func (storage *Storage) SetList(binList *bins.BinList) (bool, error) {
	if binList == nil {
		panic("binList cannot be nil")
	}

	storage.BinList = *binList
	data, err := json.MarshalIndent(storage, "", "\t")

	if err != nil {
		return false, err
	}

	db := files.NewJsonDb(fileName)
	isSuccess, err := db.Write(data)

	if err != nil {
		return false, err
	}

	return isSuccess, nil
}

func (storage *Storage) DeleteFromList(id string) {
	newList := bins.NewList()

	for _, bin := range storage.BinList.Bins {
		if bin.Id == id {
			continue
		}

		newList.Bins = append(newList.Bins, bin)
	}

	storage.SetList(newList)
}
