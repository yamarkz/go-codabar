package codabar

import (
	"fmt"
	"testing"
)

type mockCheckDigitStrategy func(seed seed) checkDigit

func (m mockCheckDigitStrategy) Apply(seed seed) checkDigit {
	return m(seed)
}

func TestNewCodabar(t *testing.T) {
	tests := []struct {
		name        string
		prefix      prefix
		body        body
		suffix      suffix
		options     []Option
		expected    codabar
		expectedErr error
	}{
		{
			name:     "No Check Digit Strategy",
			prefix:   "A",
			body:     "123456",
			suffix:   "B",
			options:  nil,
			expected: "A123456B",
		},
		{
			name:   "With Check Digit Strategy",
			prefix: "A",
			body:   "123456",
			suffix: "B",
			options: []Option{
				WithCheckDigit(CheckDigitStrategy(mockCheckDigitStrategy(func(seed seed) checkDigit {
					return checkDigit("7")
				}))),
			},
			expected: "A1234567B",
		},
		{
			name:   "With Seed Generator",
			prefix: "A",
			body:   "123456",
			suffix: "B",
			options: []Option{
				WithSeed(func(body body) (seed, error) {
					return seed(654321), nil
				}),
				WithCheckDigit(CheckDigitStrategy(mockCheckDigitStrategy(func(seed seed) checkDigit {
					return checkDigit("5")
				}))),
			},
			expected: "A1234565B",
		},
		{
			name:   "Error in Seed Generation",
			prefix: "A",
			body:   "invalid",
			suffix: "B",
			options: []Option{
				WithSeed(func(body body) (seed, error) {
					return 0, ErrInvalidArgument
				}),
				WithCheckDigit(CheckDigitStrategy(mockCheckDigitStrategy(func(seed seed) checkDigit {
					return checkDigit("0")
				}))),
			},
			expected:    "",
			expectedErr: fmt.Errorf(": %w %s", ErrInvalidArgument, "invalid"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codabar, err := NewCodabar(tt.prefix, tt.body, tt.suffix, tt.options...)
			if tt.expectedErr != nil {
				if err == nil {
					t.Fatalf("%v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("%v", err)
			}

			if codabar != tt.expected {
				t.Errorf("%v, %v", tt.expected, codabar)
			}
		})
	}
}
