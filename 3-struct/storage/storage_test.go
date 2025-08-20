package storage

import (
	"demo/go-json/bins"
	"os"
	"testing"
	"time"
)

func TestNewStorage(t *testing.T) {
	// Создаем временный файл bins.json для тестирования
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": [
				{
					"id": "test-id-1",
					"private": true,
					"createdAt": "2025-08-19T00:40:15.361Z",
					"name": "test bin 1"
				},
				{
					"id": "test-id-2",
					"private": false,
					"createdAt": "2025-08-19T00:40:18.715Z",
					"name": "test bin 2"
				}
			]
		}
	}`

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	storage := NewStorage()

	if storage == nil {
		t.Error("Expected storage to be created, got nil")
	}

	if len(storage.BinList.Bins) != 2 {
		t.Errorf("Expected 2 bins in storage, got %d", len(storage.BinList.Bins))
	}

	if storage.BinList.Bins[0].Id != "test-id-1" {
		t.Errorf("Expected first bin Id to be 'test-id-1', got %s", storage.BinList.Bins[0].Id)
	}

	if storage.BinList.Bins[1].Id != "test-id-2" {
		t.Errorf("Expected second bin Id to be 'test-id-2', got %s", storage.BinList.Bins[1].Id)
	}
}

func TestNewStorageWithEmptyFile(t *testing.T) {
	// Создаем временный пустой файл bins.json
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": []
		}
	}`

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	storage := NewStorage()

	if storage == nil {
		t.Error("Expected storage to be created, got nil")
	}

	if len(storage.BinList.Bins) != 0 {
		t.Errorf("Expected 0 bins in storage, got %d", len(storage.BinList.Bins))
	}
}

func TestSetList(t *testing.T) {
	// Создаем временный файл bins.json
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": []
		}
	}`

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	storage := NewStorage()

	// Создаем новый список бинов
	newBinList := bins.NewList()
	bin1 := bins.NewBin("new-id-1", true, time.Now(), "new bin 1")
	bin2 := bins.NewBin("new-id-2", false, time.Now(), "new bin 2")

	newBinList.AddToList(bin1)
	newBinList.AddToList(bin2)

	// Устанавливаем новый список
	success, err := storage.SetList(newBinList)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !success {
		t.Error("Expected SetList to be successful")
	}

	// Проверяем, что список был обновлен в памяти
	if len(storage.BinList.Bins) != 2 {
		t.Errorf("Expected 2 bins in storage after SetList, got %d", len(storage.BinList.Bins))
	}

	if storage.BinList.Bins[0].Id != "new-id-1" {
		t.Errorf("Expected first bin Id to be 'new-id-1', got %s", storage.BinList.Bins[0].Id)
	}

	if storage.BinList.Bins[1].Id != "new-id-2" {
		t.Errorf("Expected second bin Id to be 'new-id-2', got %s", storage.BinList.Bins[1].Id)
	}
}

func TestDeleteFromList(t *testing.T) {
	// Создаем временный файл bins.json
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": [
				{
					"id": "delete-id-1",
					"private": true,
					"createdAt": "2025-08-19T00:40:15.361Z",
					"name": "delete bin 1"
				},
				{
					"id": "keep-id-1",
					"private": false,
					"createdAt": "2025-08-19T00:40:18.715Z",
					"name": "keep bin 1"
				},
				{
					"id": "delete-id-2",
					"private": true,
					"createdAt": "2025-08-19T00:40:21.139Z",
					"name": "delete bin 2"
				}
			]
		}
	}`

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	storage := NewStorage()

	// Проверяем начальное состояние
	if len(storage.BinList.Bins) != 3 {
		t.Errorf("Expected 3 bins initially, got %d", len(storage.BinList.Bins))
	}

	// Удаляем первый элемент
	storage.DeleteFromList("delete-id-1")

	// Проверяем, что элемент был удален
	if len(storage.BinList.Bins) != 2 {
		t.Errorf("Expected 2 bins after deletion, got %d", len(storage.BinList.Bins))
	}

	// Проверяем, что оставшиеся элементы правильные
	if storage.BinList.Bins[0].Id != "keep-id-1" {
		t.Errorf("Expected first remaining bin Id to be 'keep-id-1', got %s", storage.BinList.Bins[0].Id)
	}

	if storage.BinList.Bins[1].Id != "delete-id-2" {
		t.Errorf("Expected second remaining bin Id to be 'delete-id-2', got %s", storage.BinList.Bins[1].Id)
	}

	// Удаляем еще один элемент
	storage.DeleteFromList("delete-id-2")

	// Проверяем финальное состояние
	if len(storage.BinList.Bins) != 1 {
		t.Errorf("Expected 1 bin after second deletion, got %d", len(storage.BinList.Bins))
	}

	if storage.BinList.Bins[0].Id != "keep-id-1" {
		t.Errorf("Expected remaining bin Id to be 'keep-id-1', got %s", storage.BinList.Bins[0].Id)
	}
}

func TestDeleteFromListNonExistentId(t *testing.T) {
	// Создаем временный файл bins.json
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": [
				{
					"id": "test-id-1",
					"private": true,
					"createdAt": "2025-08-19T00:40:15.361Z",
					"name": "test bin 1"
				}
			]
		}
	}`

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	storage := NewStorage()

	initialCount := len(storage.BinList.Bins)

	// Удаляем несуществующий элемент
	storage.DeleteFromList("non-existent-id")

	// Проверяем, что количество элементов не изменилось
	if len(storage.BinList.Bins) != initialCount {
		t.Errorf("Expected bin count to remain %d, got %d", initialCount, len(storage.BinList.Bins))
	}
}

func TestStorageFileNameConstant(t *testing.T) {
	if fileName != "bins.json" {
		t.Errorf("Expected fileName to be 'bins.json', got '%s'", fileName)
	}
}

func TestSetListWithNilBinList(t *testing.T) {
	// Создаем временный файл bins.json
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": []
		}
	}`

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	storage := NewStorage()

	// Тестируем установку nil списка
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when setting nil bin list")
		}
	}()

	storage.SetList(nil)
}
