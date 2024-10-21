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
	condabar, err := codabar.NewCodabar(prefix, body, suffix, codabar.WithCheckDigit(codabar.NewCheckDigitByMod11W7))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(condabar)
}
