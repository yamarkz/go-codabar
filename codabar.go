package codabar

import (
	"strconv"
)

type codabar string

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
