# go-codabar

## About
`go-codabar` generates and validates Codabar barcodes with optional check digit, written by Go.

It serves as a reference implementation to explore and verify implementation patterns, demonstrating best practices and flexible design choices.

### References
- [Codabar Specification - Morovia](https://www.morovia.com/kb/Codabar-Specification-10638.pdf)
- [NW-7 (Codabar) Barcode Technical Information - AIM Japan](https://aimjal.co.jp/gijutu/barcode/nw-7/nw-7.html)

## Highlights
Designed to emphasize Go best practices and idioms, including:

- **Idiomatic Go Code**: Follows standard naming conventions, control structures, and error handling for clarity and maintainability.
- **Simple and Effective Testing**: Utilizes table-driven tests for straightforward, extendable test cases covering various scenarios.
- **Functional Options Pattern**: Employs this pattern to allow flexible and extensible Codabar generation configurations.

## Features
- **Codabar Barcode Generation**: Create Codabar barcodes with customizable prefix, body, and suffix.
- **Check Digit Calculation**: Supports optional check digit calculation using configurable strategies (e.g., Modulus 11).
- **Flexible Configuration**: Customize behavior using functional options to specify check digit strategies, seed generation, and more.
- **Simple API**: Easy-to-use functions for common use cases.

## Provided Methods

### Codabar Generation
- `NewCodabar(prefix, body, suffix)`: Generate a standard Codabar barcode.
- `NewCodabarWithCheckDigit(prefix, body, suffix, opts...)`: Generate a Codabar barcode with an optional check digit, allowing customization through functional options.

### Check Digit Calculations
- **Modulus 11**: Provides a weighted Modulus 11 algorithm for check digit generation.
- **Custom Strategies**: Users can implement and provide their own check digit calculation logic.

## Examples

### Basic Codabar Generation
The following example demonstrates how to generate a simple Codabar barcode without a check digit:

```go
package main

import (
    "fmt"
    "github.com/yamarkz/go-codabar"
)

func main() {
    prefix, _ := codabar.NewPrefix("A")
    body, _ := codabar.NewBody("12345")
    suffix, _ := codabar.NewSuffix("B")

    code, err := codabar.NewCodabar(prefix, body, suffix)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Generated Codabar:", code)
}
```

### Codabar with Custom Check Digit

Customize the Codabar generation by providing optional check digit strategies:

```go
package main

import (
    "fmt"
    "github.com/yamarkz/go-codabar"
)

func main() {
    prefix, _ := codabar.NewPrefix("A")
    body, _ := codabar.NewBody("12345")
    suffix, _ := codabar.NewSuffix("B")

    // Define a custom check digit strategy
    checkDigitStrategy := func(s codabar.Seed) codabar.CheckDigit {
        return codabar.CheckDigit('5')
    }

    code, err := codabar.NewCodabar(
        prefix, 
        body, 
        suffix, 
        codabar.WithCheckDigit(checkDigitStrategy),
    )

    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Generated Codabar with Check Digit:", code)
}

```