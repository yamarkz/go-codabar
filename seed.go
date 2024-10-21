package codabar

import (
	"fmt"
	"strconv"
)

type seed uint64

func NewSeed(value string) (seed, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return seed(0), fmt.Errorf("seed must be a number")
	}
	return seed(parsedValue), nil
}
