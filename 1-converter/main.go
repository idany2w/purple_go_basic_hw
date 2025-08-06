package main

import "fmt"

const USD2EUR = 0.86
const USD2RUB = 0.8
const EUR2RUB = USD2RUB / USD2EUR

func main() {
	eurInput := 100.0
	rub := eurInput * EUR2RUB
	fmt.Printf("%.2f EUR is %.2f RUB\n", eurInput, rub)
}
