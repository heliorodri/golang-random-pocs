package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println("'Once upon a time' backwards is:")
	fmt.Println(stringutil.Reverse("Once upon a time"))
}
