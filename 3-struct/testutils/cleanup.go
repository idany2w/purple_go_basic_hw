package testutils

import (
	"os"
)

// TestCleanupManager управляет очисткой тестовых данных
type TestCleanupManager struct {
	tempFiles []string
	tempDirs  []string
}

// NewTestCleanupManager создает новый менеджер очистки
func NewTestCleanupManager() *TestCleanupManager {
	return &TestCleanupManager{
		tempFiles: make([]string, 0),
		tempDirs:  make([]string, 0),
	}
}

// AddTempFile добавляет временный файл для очистки
func (tcm *TestCleanupManager) AddTempFile(filename string) {
	tcm.tempFiles = append(tcm.tempFiles, filename)
}

// AddTempDir добавляет временную директорию для очистки
func (tcm *TestCleanupManager) AddTempDir(dirname string) {
	tcm.tempDirs = append(tcm.tempDirs, dirname)
}

// Cleanup очищает все временные файлы и директории
func (tcm *TestCleanupManager) Cleanup() {
	// Удаляем временные файлы
	for _, file := range tcm.tempFiles {
		if err := os.Remove(file); err != nil {
			// Логируем ошибку, но не паникуем
			// В тестах это может быть нормально, если файл уже удален
		}
	}

	// Удаляем временные директории
	for _, dir := range tcm.tempDirs {
		if err := os.RemoveAll(dir); err != nil {
			// Логируем ошибку, но не паникуем
		}
	}

	// Очищаем списки
	tcm.tempFiles = make([]string, 0)
	tcm.tempDirs = make([]string, 0)
}

// CreateTempFile создает временный файл и добавляет его в менеджер очистки
func (tcm *TestCleanupManager) CreateTempFile(content string, filename string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	tcm.AddTempFile(filename)
	return nil
}

// CreateTempDir создает временную директорию и добавляет ее в менеджер очистки
func (tcm *TestCleanupManager) CreateTempDir(dirname string) error {
	err := os.MkdirAll(dirname, 0755)
	if err != nil {
		return err
	}

	tcm.AddTempDir(dirname)
	return nil
}
