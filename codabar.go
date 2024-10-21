package codabar

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
