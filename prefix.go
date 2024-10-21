package codabar

import (
	"fmt"
	"strings"
)

type prefix string

func NewPrefix(value string) (prefix, error) {
	if len(value) != 1 || !strings.ContainsAny(value, "ABCD") {
		return prefix(""), fmt.Errorf("prefix must be a single character")
	}
	return prefix(value), nil
}
