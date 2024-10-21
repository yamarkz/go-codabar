package codabar

import "strconv"

type checkDigit string

type CheckDigitStrategy func(seed) checkDigit

// NewCheckDigitByMod11W7 は、モジュラス11 ウェイト2~7 (M11W2~7) を使用してチェックディジットを計算する。
// 計算の流れは以下の通り:
// 1. データの末尾の桁から順に、2, 3, 4, 5, 6, 7 のウェイトを適用し、それを繰り返して総和を求める。
// 2. 総和を 11 で割った余りを求める。
// 3. 11 から余りを引いた値がチェックディジットとなりる。
//   - 例外条件として、余りが 0 または 1 の場合は、チェックディジットは 0 になる。
//
// 例: シードが "12345678" の場合
// (8×2) + (7×3) + (6×4) + (5×5) + (4×6) + (3×7) + (2×2) + (1×3) = 138
// 138 % 11 = 6
// 11 - 6 = 5 → チェックディジットは "5"
func NewCheckDigitByMod11W7(seed seed) checkDigit {
	num := uint64(seed)

	var digits []uint64
	for num > 0 {
		digits = append([]uint64{num % 10}, digits...)
		num /= 10
	}

	var sum uint64
	weights := []uint64{2, 3, 4, 5, 6, 7}
	weightIndex := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * weights[weightIndex]
		weightIndex = (weightIndex + 1) % len(weights)
	}

	mod := sum % 11
	var value uint64
	if mod != 0 && mod != 1 {
		value = 11 - mod
	}

	return checkDigit(strconv.FormatUint(value, 10))
}
