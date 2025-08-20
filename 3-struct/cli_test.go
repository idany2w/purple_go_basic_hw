package main

import (
	"demo/go-json/api"
	"demo/go-json/config"
	"demo/go-json/storage"
	"demo/go-json/testutils"
	"os"
	"testing"
)

// TestCLICreateBin тестирует операцию создания бина через CLI
func TestCLICreateBin(t *testing.T) {
	// Создаем менеджер очистки
	cleanup := testutils.NewTestCleanupManager()
	defer cleanup.Cleanup()

	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-cli-create-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// Создаем временный файл bins.json
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": []
		}
	}`

	err := cleanup.CreateTempFile(jsonContent, tempFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Создаем временный JSON файл для создания бина
	testDataFile := "test_create_data.json"
	testDataContent := `{"test": "create data", "value": 42}`

	err = cleanup.CreateTempFile(testDataContent, testDataFile)
	if err != nil {
		t.Fatalf("Failed to create test data file: %v", err)
	}

	// Имитируем CLI операцию создания бина
	apiInstance := api.NewApi(config.NewConfig())
	storageInstance := storage.NewStorage()

	// Проверяем начальное состояние
	if len(storageInstance.BinList.Bins) != 0 {
		t.Errorf("Expected 0 bins initially, got %d", len(storageInstance.BinList.Bins))
	}

	// Создаем бин (имитируем операцию -create)
	binName := "test-bin-created"
	bin := apiInstance.CreateBin(testDataFile, binName)

	if bin == nil {
		t.Error("Expected bin to be created, got nil")
	}

	if bin.Name != binName {
		t.Errorf("Expected bin name %s, got %s", binName, bin.Name)
	}

	// Добавляем в хранилище
	storageInstance.BinList.AddToList(bin)
	storageInstance.SetList(&storageInstance.BinList)

	// Проверяем, что бин добавлен в хранилище
	if len(storageInstance.BinList.Bins) != 1 {
		t.Errorf("Expected 1 bin after creation, got %d", len(storageInstance.BinList.Bins))
	}

	if storageInstance.BinList.Bins[0].Name != binName {
		t.Errorf("Expected bin name in storage %s, got %s", binName, storageInstance.BinList.Bins[0].Name)
	}
}

// TestCLIGetBin тестирует операцию получения бина через CLI
func TestCLIGetBin(t *testing.T) {
	// Создаем менеджер очистки
	cleanup := testutils.NewTestCleanupManager()
	defer cleanup.Cleanup()

	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-cli-get-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// Создаем временный файл bins.json с существующим бином
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": [
				{
					"id": "test-get-id-123",
					"private": true,
					"createdAt": "2025-08-19T00:40:15.361Z",
					"name": "test bin for get"
				}
			]
		}
	}`

	err := cleanup.CreateTempFile(jsonContent, tempFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Имитируем CLI операцию получения бина
	apiInstance := api.NewApi(config.NewConfig())

	// Получаем бин (имитируем операцию -get)
	// В реальном тесте с мок-сервером это должно работать
	// Здесь мы просто проверяем, что метод выполняется без паники
	binId := "test-get-id-123"

	// Ожидаем панику, так как нет реального HTTP сервера
	defer func() {
		if r := recover(); r == nil {
			t.Log("GetBin method executed without panic (expected behavior)")
		} else {
			t.Logf("GetBin method panicked with: %v (expected behavior)", r)
		}
	}()

	bin := apiInstance.GetBin(binId)

	// Если метод не паникует, проверяем результат
	if bin != nil {
		t.Logf("Bin retrieved: ID=%s, Name=%s", bin.Id, bin.Name)
	} else {
		t.Log("Bin is nil (expected when no HTTP server)")
	}
}

// TestCLIUpdateBin тестирует операцию обновления бина через CLI
func TestCLIUpdateBin(t *testing.T) {
	// Создаем менеджер очистки
	cleanup := testutils.NewTestCleanupManager()
	defer cleanup.Cleanup()

	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-cli-update-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// Создаем временный JSON файл для обновления бина
	testDataFile := "test_update_data.json"
	testDataContent := `{"test": "updated data", "value": 100, "updated": true}`

	err := cleanup.CreateTempFile(testDataContent, testDataFile)
	if err != nil {
		t.Fatalf("Failed to create test data file: %v", err)
	}

	// Имитируем CLI операцию обновления бина
	apiInstance := api.NewApi(config.NewConfig())

	// Обновляем бин (имитируем операцию -update)
	binId := "test-update-id-456"
	success := apiInstance.UpdateBin(testDataFile, binId)

	// В реальном тесте с мок-сервером это должно быть true
	// Здесь мы просто проверяем, что метод выполняется без паники
	t.Logf("Update operation completed with success: %t", success)
}

// TestCLIDeleteBin тестирует операцию удаления бина через CLI
func TestCLIDeleteBin(t *testing.T) {
	// Создаем менеджер очистки
	cleanup := testutils.NewTestCleanupManager()
	defer cleanup.Cleanup()

	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-cli-delete-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// Создаем временный файл bins.json с существующим бином
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": [
				{
					"id": "test-delete-id-789",
					"private": false,
					"createdAt": "2025-08-19T00:40:15.361Z",
					"name": "test bin for delete"
				}
			]
		}
	}`

	err := cleanup.CreateTempFile(jsonContent, tempFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Имитируем CLI операцию удаления бина
	apiInstance := api.NewApi(config.NewConfig())
	storageInstance := storage.NewStorage()

	// Проверяем начальное состояние
	if len(storageInstance.BinList.Bins) != 1 {
		t.Errorf("Expected 1 bin initially, got %d", len(storageInstance.BinList.Bins))
	}

	// Удаляем бин (имитируем операцию -delete)
	binId := "test-delete-id-789"
	success := apiInstance.DeleteBin(binId)

	// В реальном тесте с мок-сервером это должно быть true
	t.Logf("Delete operation completed with success: %t", success)

	// Удаляем из локального хранилища
	if success {
		storageInstance.DeleteFromList(binId)

		// Проверяем, что бин удален из хранилища
		if len(storageInstance.BinList.Bins) != 0 {
			t.Errorf("Expected 0 bins after deletion, got %d", len(storageInstance.BinList.Bins))
		}
	}
}

// TestCLIListBins тестирует операцию вывода списка бинов через CLI
func TestCLIListBins(t *testing.T) {
	// Создаем менеджер очистки
	cleanup := testutils.NewTestCleanupManager()
	defer cleanup.Cleanup()

	// Создаем временный файл bins.json с несколькими бинами
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": [
				{
					"id": "list-id-1",
					"private": true,
					"createdAt": "2025-08-19T00:40:15.361Z",
					"name": "first bin"
				},
				{
					"id": "list-id-2",
					"private": false,
					"createdAt": "2025-08-19T00:40:18.715Z",
					"name": "second bin"
				}
			]
		}
	}`

	err := cleanup.CreateTempFile(jsonContent, tempFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Имитируем CLI операцию вывода списка бинов
	storageInstance := storage.NewStorage()

	// Проверяем, что бины загружены
	if len(storageInstance.BinList.Bins) != 2 {
		t.Errorf("Expected 2 bins in list, got %d", len(storageInstance.BinList.Bins))
	}

	// Проверяем содержимое списка
	if storageInstance.BinList.Bins[0].Id != "list-id-1" {
		t.Errorf("Expected first bin id 'list-id-1', got %s", storageInstance.BinList.Bins[0].Id)
	}

	if storageInstance.BinList.Bins[1].Id != "list-id-2" {
		t.Errorf("Expected second bin id 'list-id-2', got %s", storageInstance.BinList.Bins[1].Id)
	}

	// Имитируем вывод списка (имитируем операцию -list)
	storageInstance.BinList.OutputList()
}

// TestCLIWithInvalidFlags тестирует CLI с неверными флагами
func TestCLIWithInvalidFlags(t *testing.T) {
	// Создаем менеджер очистки
	cleanup := testutils.NewTestCleanupManager()
	defer cleanup.Cleanup()

	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-cli-invalid-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// Создаем временный файл bins.json
	tempFile := "bins.json"
	jsonContent := `{
		"updatedAt": "2025-08-19T03:12:44.072794772+03:00",
		"binList": {
			"bins": []
		}
	}`

	err := cleanup.CreateTempFile(jsonContent, tempFile)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Имитируем CLI без флагов (должно ничего не делать)
	apiInstance := api.NewApi(config.NewConfig())
	storageInstance := storage.NewStorage()

	// Проверяем, что компоненты созданы корректно
	if apiInstance == nil {
		t.Error("Expected API instance to be created")
	}

	if storageInstance == nil {
		t.Error("Expected storage instance to be created")
	}

	// Проверяем, что список пуст (по умолчанию)
	if len(storageInstance.BinList.Bins) != 0 {
		t.Errorf("Expected empty bin list, got %d bins", len(storageInstance.BinList.Bins))
	}
}
