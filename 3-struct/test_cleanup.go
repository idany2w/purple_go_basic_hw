package main

import (
	"demo/go-json/testutils"
	"os"
	"path/filepath"
	"testing"
)

// TestCleanupExample демонстрирует использование механизма очистки
func TestCleanupExample(t *testing.T) {
	// Создаем менеджер очистки
	cleanup := testutils.NewTestCleanupManager()

	// Гарантируем очистку после теста
	defer cleanup.Cleanup()

	// Создаем временные файлы
	err := cleanup.CreateTempFile(`{"test": "data"}`, "temp_test.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	err = cleanup.CreateTempFile(`{"another": "test"}`, "temp_another.json")
	if err != nil {
		t.Fatalf("Failed to create another temp file: %v", err)
	}

	// Создаем временную директорию
	err = cleanup.CreateTempDir("temp_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	// Проверяем, что файлы созданы
	if _, err := os.Stat("temp_test.json"); os.IsNotExist(err) {
		t.Error("Expected temp_test.json to exist")
	}

	if _, err := os.Stat("temp_another.json"); os.IsNotExist(err) {
		t.Error("Expected temp_another.json to exist")
	}

	if _, err := os.Stat("temp_test_dir"); os.IsNotExist(err) {
		t.Error("Expected temp_test_dir to exist")
	}

	// Тест завершается, cleanup.Cleanup() автоматически удалит все временные файлы
}

// TestCleanupWithSubdirectories тестирует очистку с поддиректориями
func TestCleanupWithSubdirectories(t *testing.T) {
	cleanup := testutils.NewTestCleanupManager()
	defer cleanup.Cleanup()

	// Создаем структуру директорий
	baseDir := "test_cleanup_base"
	err := cleanup.CreateTempDir(baseDir)
	if err != nil {
		t.Fatalf("Failed to create base directory: %v", err)
	}

	// Создаем поддиректории
	subDir1 := filepath.Join(baseDir, "subdir1")
	subDir2 := filepath.Join(baseDir, "subdir2")

	err = cleanup.CreateTempDir(subDir1)
	if err != nil {
		t.Fatalf("Failed to create subdir1: %v", err)
	}

	err = cleanup.CreateTempDir(subDir2)
	if err != nil {
		t.Fatalf("Failed to create subdir2: %v", err)
	}

	// Создаем файлы в поддиректориях
	err = cleanup.CreateTempFile("file1 content", filepath.Join(subDir1, "file1.txt"))
	if err != nil {
		t.Fatalf("Failed to create file1: %v", err)
	}

	err = cleanup.CreateTempFile("file2 content", filepath.Join(subDir2, "file2.txt"))
	if err != nil {
		t.Fatalf("Failed to create file2: %v", err)
	}

	// Проверяем, что структура создана
	if _, err := os.Stat(baseDir); os.IsNotExist(err) {
		t.Error("Expected base directory to exist")
	}

	if _, err := os.Stat(subDir1); os.IsNotExist(err) {
		t.Error("Expected subdir1 to exist")
	}

	if _, err := os.Stat(subDir2); os.IsNotExist(err) {
		t.Error("Expected subdir2 to exist")
	}

	// Cleanup автоматически удалит всю структуру
}

// TestCleanupMultipleTests демонстрирует использование в нескольких тестах
func TestCleanupMultipleTests(t *testing.T) {
	// Каждый тест создает свой собственный менеджер очистки
	cleanup1 := testutils.NewTestCleanupManager()
	defer cleanup1.Cleanup()

	cleanup2 := testutils.NewTestCleanupManager()
	defer cleanup2.Cleanup()

	// Создаем разные файлы для разных тестов
	err := cleanup1.CreateTempFile("test1 data", "test1.json")
	if err != nil {
		t.Fatalf("Failed to create test1 file: %v", err)
	}

	err = cleanup2.CreateTempFile("test2 data", "test2.json")
	if err != nil {
		t.Fatalf("Failed to create test2 file: %v", err)
	}

	// Проверяем, что файлы созданы
	if _, err := os.Stat("test1.json"); os.IsNotExist(err) {
		t.Error("Expected test1.json to exist")
	}

	if _, err := os.Stat("test2.json"); os.IsNotExist(err) {
		t.Error("Expected test2.json to exist")
	}

	// Каждый cleanup удалит только свои файлы
}
