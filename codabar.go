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
