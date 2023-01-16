package main

import (
	"fmt"
	"strconv"
)

func main() {
	n1 := 1
	n2 := 2

	sum := n1 + n2

	fmt.Println(strconv.Itoa(n1) + " + " + strconv.Itoa(n2))
	fmt.Println(sum)
}
