package codabar

import (
	"fmt"
	"strings"
)

type suffix string

func NewSuffix(value string) (suffix, error) {
	if len(value) != 1 || !strings.ContainsAny(value, "ABCD") {
		return suffix(""), fmt.Errorf("suffix must be a single character")
	}
	return suffix(value), nil
}
