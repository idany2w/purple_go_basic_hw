package main

import "fmt"

const USD2EUR = 0.86
const USD2RUB = 0.8
const EUR2RUB = USD2RUB / USD2EUR

func main() {
	eurInput := getUserInput()
	rub := eurInput * EUR2RUB
	fmt.Printf("%.2f EUR is %.2f RUB\n", eurInput, rub)
}

func getUserInput() float64 {
	var eurInput float64

	fmt.Println("Введите сумму в евро:")
	fmt.Scan(&eurInput)

	return eurInput
}

func convert(amount float64, fromCurrency string, toCurrency string) float64 {
}
