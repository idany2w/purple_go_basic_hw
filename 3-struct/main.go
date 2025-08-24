package main

import (
	"demo/go-json/api"
	"demo/go-json/config"
	"demo/go-json/storage"
	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	createOperation := flag.Bool("create", false, "Создание бина")
	updateOperation := flag.Bool("update", false, "Обновление бина")
	getOperation := flag.Bool("get", false, "Получение бина")
	deleteOperation := flag.Bool("delete", false, "Удаление бина")
	listOperation := flag.Bool("list", false, "Список бинов")
	id := flag.String("id", "", "Идентификатор бина")
	filePath := flag.String("file", "", "Путь к файлу")
	name := flag.String("name", "", "Название бина")

	flag.Parse()

	api := api.NewApi(config.NewConfig())
	storage := storage.NewStorage()

	if *createOperation {
		bin := api.CreateBin(*filePath, *name)
		storage.BinList.AddToList(bin)
		storage.SetList(&storage.BinList)
	} else if *getOperation {
		bin := api.GetBin(*id)
		bin.OutputBin()
	} else if *deleteOperation {
		if api.DeleteBin(*id) {
			storage.DeleteFromList(*id)
			fmt.Println("Bin successfully deleted")
		}
	} else if *listOperation {
		storage.BinList.OutputList()
	} else if *updateOperation {
		if api.UpdateBin(*filePath, *id) {
			fmt.Println("Bin successfully updated")
		}
	}
}
