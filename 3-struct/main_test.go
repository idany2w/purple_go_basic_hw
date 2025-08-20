package main

import (
	"demo/go-json/api"
	"demo/go-json/config"
	"demo/go-json/storage"
	"os"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-integration-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// Тестируем создание API
	apiInstance := api.NewApi(config.NewConfig())
	if apiInstance == nil {
		t.Error("Expected API instance to be created")
	}

	// Тестируем создание хранилища
	// Создаем временный файл bins.json для тестирования
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

	storageInstance := storage.NewStorage()
	if storageInstance == nil {
		t.Error("Expected storage instance to be created")
	}

	// Проверяем, что хранилище инициализировано с пустым списком
	if len(storageInstance.BinList.Bins) != 0 {
		t.Errorf("Expected empty bin list, got %d bins", len(storageInstance.BinList.Bins))
	}
}

func TestMainWithValidEnvironment(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-main-key"
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

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	// Создаем временный JSON файл для тестирования
	testDataFile := "test_data.json"
	testDataContent := `{"test": "data"}`

	err = os.WriteFile(testDataFile, []byte(testDataContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test data file: %v", err)
	}
	defer os.Remove(testDataFile)

	// Тестируем создание компонентов приложения
	apiInstance := api.NewApi(config.NewConfig())
	storageInstance := storage.NewStorage()

	if apiInstance == nil {
		t.Error("Expected API instance to be created")
	}

	if storageInstance == nil {
		t.Error("Expected storage instance to be created")
	}

	// Проверяем, что компоненты правильно инициализированы
	// Примечание: поле config в API является приватным, поэтому мы не можем его проверить напрямую
	// В реальном проекте стоит добавить геттер для проверки конфигурации
	t.Log("API instance created successfully")
}

func TestMainWithMissingEnvironment(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Удаляем переменную окружения
	os.Unsetenv("X_MASTER_KEY")
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// Ожидаем панику при создании конфигурации
	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic when X_MASTER_KEY is not set")
		}

		expectedMessage := "X_MASTER_KEY not found in ENV"
		if r != expectedMessage {
			t.Errorf("Expected panic message '%s', got '%v'", expectedMessage, r)
		}
	}()

	config.NewConfig()
}

// Примечание: Полные интеграционные тесты для функций main() потребуют
// мокирования флагов командной строки и HTTP запросов.
// В реальном проекте рекомендуется использовать httptest.Server для тестирования HTTP взаимодействий.

func TestMainDependencies(t *testing.T) {
	// Тест для проверки, что все зависимости доступны
	// Это помогает убедиться, что импорты работают корректно

	// Проверяем, что можем создать все необходимые типы
	var _ *api.Api
	var _ *config.Config
	var _ *storage.Storage

	// Если код компилируется, значит зависимости доступны
	t.Log("All main dependencies are available")
}

func TestMainWithNilComponents(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-nil-components"
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

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	// Тестируем создание компонентов и проверяем, что они не nil
	configInstance := config.NewConfig()
	if configInstance == nil {
		t.Error("Expected config instance to be created, got nil")
	}

	apiInstance := api.NewApi(configInstance)
	if apiInstance == nil {
		t.Error("Expected API instance to be created, got nil")
	}

	storageInstance := storage.NewStorage()
	if storageInstance == nil {
		t.Error("Expected storage instance to be created, got nil")
	}

	t.Log("All components created successfully without nil pointers")
}
