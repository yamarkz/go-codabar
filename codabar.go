package codabar

type codabar string

type prefix string
type suffix string
type body string

func NewPrefix(value string) prefix {
	return prefix(value)
}

func NewSuffix(value string) suffix {
	return suffix(value)
}

func NewBody(value string) body {
	return body(value)
}

func NewCodabar(prefix prefix, suffix suffix, body body) codabar {
	return codabar(string(prefix) + string(body) + string(suffix))
}
