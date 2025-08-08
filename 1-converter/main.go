package main

import (
	"errors"
	"fmt"
)

const USD2EUR = 0.86
const USD2RUB = 90.0
const EUR2RUB = USD2RUB / USD2EUR

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

	if fromCurrency != "USD" && fromCurrency != "EUR" && fromCurrency != "RUB" {
		fmt.Println("Неверная валюта")
		return askFromCurrency()
	}

	return fromCurrency
}

func askToCurrency(fromCurrency string) string {
	var toCurrency string

	fmt.Print("Введите валюту в которую конвертируем:")

	switch fromCurrency {
	case "USD":
		fmt.Println("EUR, RUB")
	case "EUR":
		fmt.Println("USD, RUB")
	case "RUB":
		fmt.Println("USD, EUR")
	default:
		fmt.Println("Неверная валюта")
	}

	fmt.Scan(&toCurrency)

	if toCurrency != "USD" && toCurrency != "EUR" && toCurrency != "RUB" {
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
	var result float64

	if fromCurrency == "USD" {
		switch toCurrency {
		case "EUR":
			result = amount * USD2EUR
		case "RUB":
			result = amount * USD2RUB
		}
	}

	if fromCurrency == "EUR" {
		switch toCurrency {
		case "USD":
			result = amount / USD2EUR
		case "RUB":
			result = amount * EUR2RUB
		}
	}

	if fromCurrency == "RUB" {
		switch toCurrency {
		case "USD":
			result = amount / USD2RUB
		case "EUR":
			result = amount / EUR2RUB
		}
	}

	if result == 0 {
		return 0, errors.New("invalid currency")
	}

	return result, nil
}
