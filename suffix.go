package codabar

import (
	"fmt"
	"strings"
)

type suffix string

func NewSuffix(value string) (suffix, error) {
	if len(value) != 1 || !strings.ContainsAny(value, "ABCD") {
		return suffix(""), fmt.Errorf(": %w %s", ErrInvalidArgument, value)
	}
	return suffix(value), nil
}
