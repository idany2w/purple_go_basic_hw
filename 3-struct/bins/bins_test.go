package bins

import (
	"testing"
	"time"
)

func TestNewBin(t *testing.T) {
	id := "test-id"
	private := true
	createdAt := time.Now()
	name := "test-bin"

	bin := NewBin(id, private, createdAt, name)

	if bin.Id != id {
		t.Errorf("Expected Id %s, got %s", id, bin.Id)
	}

	if bin.Private != private {
		t.Errorf("Expected Private %t, got %t", private, bin.Private)
	}

	if bin.Name != name {
		t.Errorf("Expected Name %s, got %s", name, bin.Name)
	}

	// Проверяем, что createdAt установлен в текущее время
	if bin.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set to current time")
	}
}

func TestNewList(t *testing.T) {
	binList := NewList()

	if binList == nil {
		t.Error("Expected BinList to be created, got nil")
	}

	if len(binList.Bins) != 0 {
		t.Errorf("Expected empty bins slice, got %d bins", len(binList.Bins))
	}
}

func TestAddToList(t *testing.T) {
	binList := NewList()

	bin1 := NewBin("id1", true, time.Now(), "bin1")
	bin2 := NewBin("id2", false, time.Now(), "bin2")

	binList.AddToList(bin1)

	if len(binList.Bins) != 1 {
		t.Errorf("Expected 1 bin in list, got %d", len(binList.Bins))
	}

	if binList.Bins[0].Id != "id1" {
		t.Errorf("Expected first bin Id to be 'id1', got %s", binList.Bins[0].Id)
	}

	binList.AddToList(bin2)

	if len(binList.Bins) != 2 {
		t.Errorf("Expected 2 bins in list, got %d", len(binList.Bins))
	}

	if binList.Bins[1].Id != "id2" {
		t.Errorf("Expected second bin Id to be 'id2', got %s", binList.Bins[1].Id)
	}
}

func TestAddToListWithNilBin(t *testing.T) {
	binList := NewList()

	// Тестируем добавление nil бина
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when adding nil bin")
		}
	}()

	binList.AddToList(nil)
}

func TestOutputBin(t *testing.T) {
	// Этот тест проверяет, что метод не паникует
	// В реальном проекте можно было бы перехватывать stdout
	bin := NewBin("test-id", true, time.Now(), "test-name")

	// Просто проверяем, что метод выполняется без ошибок
	bin.OutputBin()
}

func TestOutputList(t *testing.T) {
	binList := NewList()

	bin1 := NewBin("id1", true, time.Now(), "bin1")
	bin2 := NewBin("id2", false, time.Now(), "bin2")

	binList.AddToList(bin1)
	binList.AddToList(bin2)

	// Просто проверяем, что метод выполняется без ошибок
	binList.OutputList()
}

func TestEmptyListOutput(t *testing.T) {
	binList := NewList()

	// Проверяем, что вывод пустого списка не вызывает ошибок
	binList.OutputList()
}
