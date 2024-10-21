package codabar

import (
	"fmt"
	"strconv"
)

type body string

func NewBody(value string) (body, error) {
	if _, err := strconv.ParseUint(value, 10, 64); err != nil {
		return body(""), fmt.Errorf("body must be a number")
	}
	return body(value), nil
}
