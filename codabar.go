package codabar

import (
	"fmt"
	"strconv"
	"strings"
)

type codabar string

type prefix string
type suffix string
type body string

func NewPrefix(value string) (prefix, error) {
	if len(value) != 1 || !strings.ContainsAny(value, "ABCD") {
		return prefix(""), fmt.Errorf("prefix must be a single character")
	}
	return prefix(value), nil
}

func NewSuffix(value string) (suffix, error) {
	if len(value) != 1 || !strings.ContainsAny(value, "ABCD") {
		return suffix(""), fmt.Errorf("suffix must be a single character")
	}
	return suffix(value), nil
}

func NewBody(value string) (body, error) {
	if _, err := strconv.ParseUint(value, 10, 64); err != nil {
		return body(""), fmt.Errorf("body must be a number")
	}
	return body(value), nil
}

func NewCodabar(prefix prefix, suffix suffix, body body) codabar {
	return codabar(string(prefix) + string(body) + string(suffix))
}

func NewCodabarWithCheckDigit(prefix prefix, suffix suffix, body body, strategy CheckDigitStrategy) (codabar, error) {
	seed, err := NewSeed(string(body))
	if err != nil {
		return codabar(""), err
	}
	checkDigit := strategy(seed)
	return codabar(string(prefix) + string(body) + string(checkDigit) + string(suffix)), nil
}

type checkDigit string
type seed uint64

func NewSeed(value string) (seed, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return seed(0), fmt.Errorf("seed must be a number")
	}
	return seed(parsedValue), nil
}

type CheckDigitStrategy func(seed) checkDigit

func NewCheckDigitByMod11W7(seed seed) checkDigit {
	num := uint64(seed)

	var digits []uint64
	for num > 0 {
		digits = append([]uint64{num % 10}, digits...)
		num /= 10
	}

	var sum uint64
	weights := []uint64{2, 3, 4, 5, 6, 7}
	weightIndex := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * weights[weightIndex]
		weightIndex = (weightIndex + 1) % len(weights)
	}

	mod := sum % 11
	var value uint64
	if mod != 0 && mod != 1 {
		value = 11 - mod
	}

	return checkDigit(strconv.FormatUint(value, 10))
}
