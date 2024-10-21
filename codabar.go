package codabar

import (
	"errors"
	"fmt"
)

type codabar string

var ErrInvalidArgument = errors.New("invalid argument")

type Option func(*codabarConfig)

type codabarConfig struct {
	checkDigitStrategy CheckDigitStrategy
	seedGenerator      func(body) (seed, error)
}

func WithCheckDigit(strategy CheckDigitStrategy) Option {
	return func(cfg *codabarConfig) {
		cfg.checkDigitStrategy = strategy
	}
}

func WithSeed(generator func(body) (seed, error)) Option {
	return func(cfg *codabarConfig) {
		cfg.seedGenerator = generator
	}
}

func NewCodabar(prefix prefix, b body, suffix suffix, opts ...Option) (codabar, error) {
	cfg := &codabarConfig{
		seedGenerator: func(b body) (seed, error) {
			return NewSeed(string(b))
		},
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.checkDigitStrategy != nil {
		seed, err := cfg.seedGenerator(b)
		if err != nil {
			return codabar(""), fmt.Errorf(": %w %s", ErrInvalidArgument, b)
		}
		checkDigit := cfg.checkDigitStrategy(seed)
		return codabar(string(prefix) + string(b) + string(checkDigit) + string(suffix)), nil
	}

	return codabar(string(prefix) + string(b) + string(suffix)), nil
}
