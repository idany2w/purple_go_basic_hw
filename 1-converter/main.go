package main

import (
	"errors"
	"fmt"
)

const USD2EUR = 0.86
const USD2RUB = 90.0
const EUR2RUB = USD2RUB / USD2EUR

// Константы для валют
const (
	USD = "USD"
	EUR = "EUR"
	RUB = "RUB"
)

// Map для конвертации валют
var conversionRates = map[string]map[string]float64{
	USD: {
		EUR: USD2EUR,
		RUB: USD2RUB,
	},
	EUR: {
		USD: 1 / USD2EUR,
		RUB: EUR2RUB,
	},
	RUB: {
		USD: 1 / USD2RUB,
		EUR: 1 / EUR2RUB,
	},
}

// Map для отображения доступных валют
var availableCurrencies = map[string][]string{
	USD: {EUR, RUB},
	EUR: {USD, RUB},
	RUB: {USD, EUR},
}

func main() {
	for {
		processConverting()

		fmt.Println("Хотите попробовать еще раз? (y/n)")
		var continueInput string
		fmt.Scan(&continueInput)

		if continueInput != "y" && continueInput != "Y" {
			break
		}
	}
}

func processConverting() {
	amount, fromCurrency, toCurrency := getUserInput()
	result, err := convert(amount, fromCurrency, toCurrency)

	if err != nil {
		fmt.Println("Неверная валюта")
		return
	}

	fmt.Printf("%.2f %s is %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

func getUserInput() (float64, string, string) {
	fromCurrency := askFromCurrency()
	amount := askAmount()
	toCurrency := askToCurrency(fromCurrency)
	return amount, fromCurrency, toCurrency
}

func askFromCurrency() string {
	var fromCurrency string
	fmt.Println("Введите валюту из которой конвертируем: (USD, EUR, RUB)")
	fmt.Scan(&fromCurrency)

	if _, exists := availableCurrencies[fromCurrency]; !exists {
		fmt.Println("Неверная валюта")
		return askFromCurrency()
	}

	return fromCurrency
}

func askToCurrency(fromCurrency string) string {
	var toCurrency string

	fmt.Print("Введите валюту в которую конвертируем:")

	if available, exists := availableCurrencies[fromCurrency]; exists {
		fmt.Println(available[0] + ", " + available[1])
	} else {
		fmt.Println("Неверная валюта")
	}

	fmt.Scan(&toCurrency)

	if _, exists := availableCurrencies[toCurrency]; !exists {
		fmt.Println("Неверная валюта")
		return askToCurrency(fromCurrency)
	}

	if fromCurrency == toCurrency {
		fmt.Println("Валюты не могут быть одинаковыми")
		return askToCurrency(fromCurrency)
	}

	return toCurrency
}

func askAmount() float64 {
	var amount float64
	fmt.Println("Введите сумму:")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Сумма должна быть больше 0")
		return askAmount()
	}

	return amount
}

func convert(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	if rates, exists := conversionRates[fromCurrency]; exists {
		if rate, rateExists := rates[toCurrency]; rateExists {
			return amount * rate, nil
		}
	}

	return 0, errors.New("invalid currency")
}
