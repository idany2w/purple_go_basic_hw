package files

import (
	"os"
	"testing"
)

func TestJsonDbWithRealTestFile(t *testing.T) {
	// Используем реальный тестовый файл из корня проекта
	testFile := "../test_data.json"

	// Проверяем, что файл существует
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		t.Skipf("Test file %s does not exist, skipping test", testFile)
	}

	jsonDb := NewJsonDb(testFile)

	if jsonDb == nil {
		t.Error("Expected JsonDb to be created, got nil")
	}

	if jsonDb.filename != testFile {
		t.Errorf("Expected filename %s, got %s", testFile, jsonDb.filename)
	}

	// Читаем содержимое файла
	data, err := jsonDb.Read()

	if err != nil {
		t.Errorf("Expected no error when reading test file, got %v", err)
	}

	if len(data) == 0 {
		t.Error("Expected non-empty data from test file")
	}

	// Проверяем, что это валидный JSON
	expectedStart := `{`
	if string(data[:1]) != expectedStart {
		t.Errorf("Expected JSON to start with '%s', got '%s'", expectedStart, string(data[:1]))
	}
}

func TestJsonDbWriteAndReadCycle(t *testing.T) {
	tempFile := "test_cycle.json"
	testContent := `{"cycle": "test", "data": 123}`

	jsonDb := NewJsonDb(tempFile)

	// Записываем данные
	success, err := jsonDb.Write([]byte(testContent))
	if err != nil {
		t.Fatalf("Failed to write test data: %v", err)
	}

	if !success {
		t.Error("Expected write to be successful")
	}

	// Читаем данные обратно
	readData, err := jsonDb.Read()
	if err != nil {
		t.Fatalf("Failed to read test data: %v", err)
	}

	// Проверяем, что данные совпадают
	if string(readData) != testContent {
		t.Errorf("Expected read data to match written data. Expected: %s, Got: %s", testContent, string(readData))
	}

	// Очищаем
	os.Remove(tempFile)
}

func TestJsonDbMultipleWrites(t *testing.T) {
	tempFile := "test_multiple.json"

	jsonDb := NewJsonDb(tempFile)

	// Первая запись
	content1 := `{"first": "write"}`
	success, err := jsonDb.Write([]byte(content1))
	if err != nil {
		t.Fatalf("Failed first write: %v", err)
	}

	if !success {
		t.Error("Expected first write to be successful")
	}

	// Вторая запись (должна перезаписать первую)
	content2 := `{"second": "write"}`
	success, err = jsonDb.Write([]byte(content2))
	if err != nil {
		t.Fatalf("Failed second write: %v", err)
	}

	if !success {
		t.Error("Expected second write to be successful")
	}

	// Читаем и проверяем, что получили вторую запись
	readData, err := jsonDb.Read()
	if err != nil {
		t.Fatalf("Failed to read after multiple writes: %v", err)
	}

	if string(readData) != content2 {
		t.Errorf("Expected to read second write content. Expected: %s, Got: %s", content2, string(readData))
	}

	// Очищаем
	os.Remove(tempFile)
}
