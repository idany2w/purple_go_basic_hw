package main

import "fmt"

const USD2EUR = 0.86
const USD2RUB = 0.8

func main() {
	eurInput := 100.0
	usd := eurInput / USD2EUR
	rub := usd * USD2RUB
	fmt.Printf("%.2f EUR is %.2f USD and %.2f RUB\n", eurInput, usd, rub)
}
