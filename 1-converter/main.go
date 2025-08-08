package main

import "fmt"

const USD2EUR = 0.86
const USD2RUB = 0.8
const EUR2RUB = USD2RUB / USD2EUR

func main() {
	userInput := getUserInput()
	rub := convert(userInput, "EUR", "RUB")
	fmt.Printf("%.2f EUR is %.2f RUB\n", userInput, rub)
}

func getUserInput() float64 {
	var eurInput float64

	fmt.Println("Введите сумму:")
	fmt.Scan(&eurInput)

	return eurInput
}

func convert(amount float64, fromCurrency string, toCurrency string) float64 {
	return amount
}
