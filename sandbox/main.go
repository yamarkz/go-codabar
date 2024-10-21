package main

import (
	"fmt"

	"github.com/yamarkz/go-codabar"
)

func main() {
	prefix, err := codabar.NewPrefix("A")
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := codabar.NewBody("12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	suffix, err := codabar.NewSuffix("B")
	if err != nil {
		fmt.Println(err)
		return
	}
	condabar := codabar.NewCodabar(prefix, body, suffix)
	fmt.Println(condabar)
}
