package codabar

import (
	"strconv"
	"testing"
)

func TestNewCheckDigitByMod11W7(t *testing.T) {
	tests := []struct {
		seed     seed
		expected checkDigit
	}{
		{12345, "5"},
		{67890, "2"},
		{123456, "0"},
		{111111, "6"},
		{222222, "1"},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatUint(uint64(tt.seed), 10), func(t *testing.T) {
			result := NewCheckDigitByMod11W7(tt.seed)
			if result != tt.expected {
				t.Errorf("NewCheckDigitByMod11W7(%d) = %s; want %s", tt.seed, result, tt.expected)
			}
		})
	}
}
