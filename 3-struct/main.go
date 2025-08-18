package main

import (
	"demo/struct/src/bins"
	"demo/struct/src/file"
	"demo/struct/src/storage"
	"fmt"
)

func main() {
	fp, err := file.NewFileProvider("data.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	storage := storage.NewStorage(fp)

	bin := bins.NewBin("1", true, "test")
	storage.AddBin(*bin)

	err = storage.SaveBinList()

	if err != nil {
		fmt.Println(err)
		return
	}
}
