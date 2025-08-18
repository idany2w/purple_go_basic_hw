package storage

import (
	"demo/struct/src/bins"
	"encoding/json"
)

type StorageProviderInterface interface {
	Read() ([]byte, error)
	Save(data []byte) error
}

type Storage struct {
	bl       bins.BinList
	provider StorageProviderInterface
}

func NewStorage(provider StorageProviderInterface) *Storage {
	return &Storage{
		bl:       bins.BinList{},
		provider: provider,
	}
}

func (s *Storage) Read() (bins.BinList, error) {
	var bl bins.BinList
	data, err := s.provider.Read()

	if err != nil {
		return bl, err
	}

	err = json.Unmarshal(data, &bl)

	if err != nil {
		return bl, err
	}

	return bl, nil
}

func (s *Storage) SaveBinList() error {
	data, err := json.Marshal(s.bl)

	if err != nil {
		return err
	}

	return s.provider.Save(data)
}

func (s *Storage) AddBin(bin bins.Bin) {
	s.bl.Bins = append(s.bl.Bins, bin)
}
