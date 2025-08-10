package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var allowedOperations = [3]string{
	"avg", // среднее значение
	"sum", // сумма
	"med", // медиана
}

func main() {
	operation, err := askOperation()
	if err != nil {
		fmt.Println(err)
		return
	}

	numbers, err := askNumbers()
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := doOperation(operation, numbers)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}

func askOperation() (string, error) {
	fmt.Println("Введите операцию (avg, sum, med):")
	var operation string
	fmt.Scan(&operation)

	if !isValidOperation(operation) {
		return "", errors.New("неподдерживаемая операция. Поддерживаемые операции: avg, sum, med")
	}

	return operation, nil
}

func isValidOperation(operation string) bool {
	for _, op := range allowedOperations {
		if op == operation {
			return true
		}
	}
	return false
}

func askNumbers() ([]float64, error) {
	fmt.Println("Введите числа через запятую:")
	var numbersString string
	fmt.Scan(&numbersString)

	numbersArray := strings.Split(numbersString, ",")

	if len(numbersArray) == 0 {
		return []float64{}, errors.New("не удалось разделить числа")
	}

	if strings.TrimSpace(numbersString) == "" {
		return []float64{}, errors.New("введите хотя бы одно число")
	}

	numbers := make([]float64, 0, len(numbersArray))

	for i, v := range numbersArray {
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			continue
		}

		num, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return []float64{}, fmt.Errorf("некорректное число '%s' на позиции %d", trimmed, i+1)
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		return []float64{}, errors.New("не найдено корректных чисел для обработки")
	}

	return numbers, nil
}

// Map для отображения операций на соответствующие функции
var operationFunctions = map[string]func([]float64) float64{
	"avg": avg,
	"sum": sum,
	"med": med,
}

func doOperation(operation string, numbers []float64) (float64, error) {
	operationFunc, exists := operationFunctions[operation]
	if !exists {
		return 0, errors.New("неизвестная операция")
	}

	result := operationFunc(numbers)
	return result, nil
}

func sum(numbers []float64) float64 {
	sum := 0.0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func avg(numbers []float64) float64 {
	return sum(numbers) / float64(len(numbers))
}

func med(numbers []float64) float64 {
	sort.Float64s(numbers)
	mid := len(numbers) / 2

	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2
	}

	return numbers[mid]
}
