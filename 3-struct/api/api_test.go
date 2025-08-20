package api

import (
	"demo/go-json/config"
	"os"
	"testing"
)

func TestNewApi(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-api-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	config := config.NewConfig()
	api := NewApi(config)

	if api == nil {
		t.Error("Expected API to be created, got nil")
	}

	if api.config.XMasterKey != testKey {
		t.Errorf("Expected XMasterKey %s, got %s", testKey, api.config.XMasterKey)
	}
}

func TestApiUrlConstant(t *testing.T) {
	expectedUrl := "https://api.jsonbin.io/v3/b"
	if apiUrl != expectedUrl {
		t.Errorf("Expected apiUrl to be '%s', got '%s'", expectedUrl, apiUrl)
	}
}

func TestGetHttpClient(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-master-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	config := config.NewConfig()
	api := NewApi(config)

	// Тестируем GET запрос
	client, req := api.getHttpClient("GET", "https://example.com", nil)

	if client == nil {
		t.Error("Expected HTTP client to be created, got nil")
	}

	if req == nil {
		t.Error("Expected HTTP request to be created, got nil")
	}

	if req.Method != "GET" {
		t.Errorf("Expected method 'GET', got '%s'", req.Method)
	}

	if req.URL.String() != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got '%s'", req.URL.String())
	}

	// Проверяем заголовки
	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', got '%s'", contentType)
	}

	masterKey := req.Header.Get("X-Master-Key")
	if masterKey != testKey {
		t.Errorf("Expected X-Master-Key '%s', got '%s'", testKey, masterKey)
	}
}

func TestGetHttpClientWithData(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-master-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	config := config.NewConfig()
	api := NewApi(config)

	testData := []byte(`{"test": "data"}`)
	client, req := api.getHttpClient("POST", "https://example.com", testData)

	if client == nil {
		t.Error("Expected HTTP client to be created, got nil")
	}

	if req == nil {
		t.Error("Expected HTTP request to be created, got nil")
	}

	if req.Method != "POST" {
		t.Errorf("Expected method 'POST', got '%s'", req.Method)
	}

	// Проверяем, что тело запроса содержит данные
	if req.Body == nil {
		t.Error("Expected request body to be set")
	}
}

func TestCreateBinResponseStructure(t *testing.T) {
	// Проверяем, что структура CreateBinResponse существует и имеет правильное поле
	var response CreateBinResponse

	// Это просто проверка компиляции - структура должна существовать
	_ = response.MetaData
}

// Примечание: Полные интеграционные тесты для CreateBin, UpdateBin, GetBin и DeleteBin
// потребуют мокирования HTTP клиента или использования тестового сервера.
// В реальном проекте рекомендуется использовать httptest.Server или библиотеки для мокирования.

func TestApiWithNilConfig(t *testing.T) {
	// Тест для проверки поведения с nil конфигурацией
	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic when config is nil")
		}

		expectedMessage := "config cannot be nil"
		if r != expectedMessage {
			t.Errorf("Expected panic message '%s', got '%v'", expectedMessage, r)
		}
	}()

	// Вызываем NewApi с nil конфигурацией
	NewApi(nil)
}
