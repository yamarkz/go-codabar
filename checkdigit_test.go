package codabar

import (
	"strconv"
	"testing"
)

func TestNewCheckDigitByMod10W21Division(t *testing.T) {
	tests := []struct {
		seed     seed
		expected checkDigit
	}{
		{12345678, "2"},
		{87654321, "6"},
		{1234, "4"},
		{5678, "8"},
		{111111, "1"},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatUint(uint64(tt.seed), 10), func(t *testing.T) {
			result := NewCheckDigitByMod10W21Division(tt.seed)
			if result != tt.expected {
				t.Errorf("NewCheckDigitByMod10W21Division(%d) = %s; want %s", tt.seed, result, tt.expected)
			}
		})
	}
}

func TestNewCheckDigitByMod10W21Bulk(t *testing.T) {
	tests := []struct {
		seed     seed
		expected checkDigit
	}{
		{12345678, "4"},
		{87654321, "8"},
		{1234, "4"},
		{5678, "0"},
		{111111, "1"},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatUint(uint64(tt.seed), 10), func(t *testing.T) {
			result := NewCheckDigitByMod10W21Bulk(tt.seed)
			if result != tt.expected {
				t.Errorf("NewCheckDigitByMod10W21Bulk(%d) = %s; want %s", tt.seed, result, tt.expected)
			}
		})
	}
}

func TestNewCheckDigitByMod10W31(t *testing.T) {
	tests := []struct {
		seed     seed
		expected checkDigit
	}{
		{12345, "7"},
		{67890, "2"},
		{11111, "9"},
		{22222, "8"},
		{33333, "7"},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatUint(uint64(tt.seed), 10), func(t *testing.T) {
			result := NewCheckDigitByMod10W31(tt.seed)
			if result != tt.expected {
				t.Errorf("NewCheckDigitByMod10W31(%d) = %s; want %s", tt.seed, result, tt.expected)
			}
		})
	}
}

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

func TestNewCheckDigitByMod11W10(t *testing.T) {
	tests := []struct {
		seed     seed
		expected checkDigit
	}{
		{12345678, "1"},
		{87654321, "5"},
		{111111, "1"},
		{222222, "2"},
		{333333, "3"},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatUint(uint64(tt.seed), 10), func(t *testing.T) {
			result := NewCheckDigitByMod11W10(tt.seed)
			if result != tt.expected {
				t.Errorf("NewCheckDigitByMod11W10(%d) = %s; want %s", tt.seed, result, tt.expected)
			}
		})
	}
}

func TestNewCheckDigitBySevenCheck(t *testing.T) {
	tests := []struct {
		seed     seed
		expected checkDigit
	}{
		{12345, "3"},
		{100, "5"},
		{77, "0"},
		{49, "0"},
		{65432, "4"},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatUint(uint64(tt.seed), 10), func(t *testing.T) {
			result := NewCheckDigitBySevenCheck(tt.seed)
			if result != tt.expected {
				t.Errorf("NewCheckDigitBySevenCheck(%d) = %s; want %s", tt.seed, result, tt.expected)
			}
		})
	}
}

func TestNewCheckDigitByNineCheck(t *testing.T) {
	tests := []struct {
		seed     seed
		expected checkDigit
	}{
		{12345, "3"},
		{100, "8"},
		{81, "0"},
		{72, "0"},
		{65432, "7"},
	}

	for _, tt := range tests {
		t.Run(strconv.FormatUint(uint64(tt.seed), 10), func(t *testing.T) {
			result := NewCheckDigitByNineCheck(tt.seed)
			if result != tt.expected {
				t.Errorf("NewCheckDigitByNineCheck(%d) = %s; want %s", tt.seed, result, tt.expected)
			}
		})
	}
}
