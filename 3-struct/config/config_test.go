package config

import (
	"os"
	"testing"
)

func TestNewConfigWithValidEnvVar(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-master-key-123"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	config := NewConfig()

	if config == nil {
		t.Error("Expected config to be created, got nil")
	}

	if config.XMasterKey != testKey {
		t.Errorf("Expected XMasterKey %s, got %s", testKey, config.XMasterKey)
	}
}

func TestNewConfigWithEmptyEnvVar(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Удаляем переменную окружения
	os.Unsetenv("X_MASTER_KEY")
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// Ожидаем панику
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

	NewConfig()
}

func TestNewConfigWithWhitespaceEnvVar(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем значение с пробелами
	os.Setenv("X_MASTER_KEY", "   ")
	defer os.Setenv("X_MASTER_KEY", originalKey)

	// В текущей реализации os.Getenv возвращает строку с пробелами,
	// которая не является пустой, поэтому паника не должна происходить
	config := NewConfig()

	if config == nil {
		t.Error("Expected config to be created, got nil")
	}

	if config.XMasterKey != "   " {
		t.Errorf("Expected XMasterKey to be '   ', got '%s'", config.XMasterKey)
	}
}

func TestConfigStructure(t *testing.T) {
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "another-test-key"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	config := NewConfig()

	// Проверяем, что структура Config имеет правильное поле
	if config.XMasterKey == "" {
		t.Error("Expected XMasterKey to be set")
	}
}

func TestConfigWithNilPointer(t *testing.T) {
	// Тест для проверки, что конфигурация не возвращает nil
	// Сохраняем оригинальное значение переменной окружения
	originalKey := os.Getenv("X_MASTER_KEY")

	// Устанавливаем тестовое значение
	testKey := "test-nil-check"
	os.Setenv("X_MASTER_KEY", testKey)
	defer os.Setenv("X_MASTER_KEY", originalKey)

	config := NewConfig()

	if config == nil {
		t.Error("Expected config to be created, got nil")
	}

	// Проверяем, что можем безопасно обращаться к полям
	if config.XMasterKey != testKey {
		t.Errorf("Expected XMasterKey %s, got %s", testKey, config.XMasterKey)
	}
}
