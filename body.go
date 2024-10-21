package codabar

import (
	"fmt"
	"strconv"
)

type body string

func NewBody(value string) (body, error) {
	if _, err := strconv.ParseUint(value, 10, 64); err != nil {
		return body(""), fmt.Errorf(": %w %s", ErrInvalidArgument, value)
	}
	return body(value), nil
}
