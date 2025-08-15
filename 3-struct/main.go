package main

import (
	"demo/struct/src/bins"
	"fmt"
)

func main() {
	bin := bins.NewBin("id", false, "name")
	fmt.Println(*bin)
}
