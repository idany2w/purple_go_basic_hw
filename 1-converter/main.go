package main

import "fmt"

const USD2EUR = 0.86
const USD2RUB = 0.8

func main() {
	eurInput := 100.0

	usd := (1.0 / USD2EUR) * eurInput

	rub := usd * USD2RUB

	fmt.Println(rub)
}
