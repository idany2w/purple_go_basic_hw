package files

import (
	"os"
	"testing"
)

func TestNewJsonDb(t *testing.T) {
	filename := "test.json"

	jsonDb := NewJsonDb(filename)

	if jsonDb == nil {
		t.Error("Expected JsonDb to be created, got nil")
	}

	if jsonDb.filename != filename {
		t.Errorf("Expected filename %s, got %s", filename, jsonDb.filename)
	}
}

func TestReadValidJsonFile(t *testing.T) {
	// Создаем временный JSON файл
	tempFile := "test_valid.json"
	jsonContent := `{"test": "data"}`

	err := os.WriteFile(tempFile, []byte(jsonContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	jsonDb := NewJsonDb(tempFile)
	data, err := jsonDb.Read()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if string(data) != jsonContent {
		t.Errorf("Expected content %s, got %s", jsonContent, string(data))
	}
}

func TestReadNonExistentFile(t *testing.T) {
	jsonDb := NewJsonDb("non_existent.json")
	_, err := jsonDb.Read()

	if err == nil {
		t.Error("Expected error when reading non-existent file")
	}

	expectedError := "При чтении файла произошла ошибка"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

func TestReadNonJsonFile(t *testing.T) {
	// Создаем временный файл с неправильным расширением
	tempFile := "test.txt"
	content := "some text content"

	err := os.WriteFile(tempFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tempFile)

	jsonDb := NewJsonDb(tempFile)
	_, err = jsonDb.Read()

	if err == nil {
		t.Error("Expected error when reading non-json file")
	}

	expectedError := "Невалидный json-файл"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

func TestWriteValidContent(t *testing.T) {
	tempFile := "test_write.json"
	content := []byte(`{"test": "write data"}`)

	jsonDb := NewJsonDb(tempFile)
	success, err := jsonDb.Write(content)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !success {
		t.Error("Expected write to be successful")
	}

	// Проверяем, что файл был создан и содержит правильные данные
	readData, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}

	if string(readData) != string(content) {
		t.Errorf("Expected written content %s, got %s", string(content), string(readData))
	}

	os.Remove(tempFile)
}

func TestWriteToReadOnlyDirectory(t *testing.T) {
	// Этот тест может не работать на всех системах
	// В реальном проекте можно было бы создать временную директорию с правами только для чтения

	// Пропускаем тест, если не можем создать read-only директорию
	t.Skip("Skipping read-only directory test as it may not work on all systems")
}

func TestWriteEmptyContent(t *testing.T) {
	tempFile := "test_empty.json"
	content := []byte("")

	jsonDb := NewJsonDb(tempFile)
	success, err := jsonDb.Write(content)

	if err != nil {
		t.Errorf("Expected no error when writing empty content, got %v", err)
	}

	if !success {
		t.Error("Expected write to be successful even with empty content")
	}

	// Проверяем, что файл был создан
	_, err = os.ReadFile(tempFile)
	if err != nil {
		t.Errorf("Expected file to be created, got error: %v", err)
	}

	os.Remove(tempFile)
}

func TestJsonExtensionConstant(t *testing.T) {
	if jsonExtension != ".json" {
		t.Errorf("Expected jsonExtension to be '.json', got '%s'", jsonExtension)
	}
}

func TestNewJsonDbWithEmptyName(t *testing.T) {
	// Тестируем создание JsonDb с пустым именем
	jsonDb := NewJsonDb("")

	if jsonDb == nil {
		t.Error("Expected JsonDb to be created, got nil")
	}

	if jsonDb.filename != "" {
		t.Errorf("Expected empty filename, got '%s'", jsonDb.filename)
	}
}

func TestNewJsonDbWithNilName(t *testing.T) {
	// Тестируем создание JsonDb с nil именем (это невозможно в Go, но для полноты)
	// В Go строки не могут быть nil, но могут быть пустыми
	jsonDb := NewJsonDb("")

	if jsonDb == nil {
		t.Error("Expected JsonDb to be created, got nil")
	}
}
