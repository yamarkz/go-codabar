package codabar

import (
	"fmt"
	"strconv"
)

type seed uint64

func NewSeed(value string) (seed, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return seed(0), fmt.Errorf(": %w %s", ErrInvalidArgument, value)
	}
	return seed(parsedValue), nil
}
