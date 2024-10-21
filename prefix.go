package codabar

import (
	"fmt"
	"strings"
)

type prefix string

func NewPrefix(value string) (prefix, error) {
	if len(value) != 1 || !strings.ContainsAny(value, "ABCD") {
		return prefix(""), fmt.Errorf(": %w %s", ErrInvalidArgument, value)
	}
	return prefix(value), nil
}
