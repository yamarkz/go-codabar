package codabar

import "strconv"

type checkDigit string

type CheckDigitStrategy func(seed) checkDigit

// NewCheckDigitByMod10W21Division は、モジュラス10 ウェイト2・1分割 (M10W21) を使用してチェックディジットを計算する。
// 計算の流れは以下の通り:
// 1. データの末尾の桁から順に、2 と 1 を交互に掛けて総和を求める。
// 2. 2桁になった場合は各桁を分けて加算する。
// 3. 総和を 10 で割った余りを求める。
// 4. 10 から余りを引いた値がチェックディジットとなる。
//   - 余りが 0 の場合はチェックディジットも 0 になる。
//
// 例: シードが "12345678" の場合
// (8×2)+(7×1)+(6×2)+(5×1)+(4×2)+(3×1)+(2×2)+(1×1) = 16+7+12+5+8+3+4+1
// (1+6) + 7 + (1+2) + 5 + 8 + 3 + 4 + 1 = 38
// 38 % 10 = 8
// 10 - 8 = 2 → チェックディジットは "2"
func NewCheckDigitByMod10W21Division(seed seed) checkDigit {
	num := uint64(seed)

	var digits []uint64
	for num > 0 {
		digits = append([]uint64{num % 10}, digits...)
		num /= 10
	}

	var sum uint64
	weights := []uint64{2, 1}
	weightIndex := 0
	for i := len(digits) - 1; i >= 0; i-- {
		product := digits[i] * weights[weightIndex]

		if product >= 10 {
			sum += product/10 + product%10
		} else {
			sum += product
		}

		weightIndex = (weightIndex + 1) % len(weights)
	}

	mod := sum % 10
	var value uint64
	if mod != 0 {
		value = 10 - mod
	}

	return checkDigit(strconv.FormatUint(value, 10))
}

// NewCheckDigitByMod10W21Bulk は、モジュラス10 ウェイト2・1一括 (M10W21) を使用してチェックディジットを計算する。
// 計算の流れは以下の通り:
// 1. データの末尾の桁から順に、2 と 1 を交互に掛けて総和を求める。
// 2. 総和を 10 で割った余りを求める。
// 3. 10 から余りを引いた値がチェックディジットとなる。
//   - 余りが 0 の場合はチェックディジットも 0 になる。
//
// 例: シードが "12345678" の場合
// (8×2)+(7×1)+(6×2)+(5×1)+(4×2)+(3×1)+(2×2)+(1×1) = 16+7+12+5+8+3+4+1 = 56
// 56 % 10 = 6
// 10 - 6 = 4 → チェックディジットは "4"
func NewCheckDigitByMod10W21Bulk(seed seed) checkDigit {
	num := uint64(seed)

	var digits []uint64
	for num > 0 {
		digits = append([]uint64{num % 10}, digits...)
		num /= 10
	}

	var sum uint64
	weights := []uint64{2, 1}
	weightIndex := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * weights[weightIndex]
		weightIndex = (weightIndex + 1) % len(weights)
	}

	mod := sum % 10
	var value uint64
	if mod != 0 {
		value = 10 - mod
	}

	return checkDigit(strconv.FormatUint(value, 10))
}

// NewCheckDigitByMod10W31 は、モジュラス10 ウェイト3・1 (M10W31) を使用してチェックディジットを計算する。
// 計算の流れは以下の通り:
// 1. データの末尾の桁から順に、3 と 1 を交互に掛けて総和を求める。
// 2. 総和を 10 で割った余りを求める。
// 3. 10 から余りを引いた値がチェックディジットとなる。
//   - 余りが 0 の場合はチェックディジットも 0 になる。
//
// 例: シードが "12345" の場合
// (5×3)+(4×1)+(3×3)+(2×1)+(1×3) = 33
// 33 % 10 = 3
// 10 - 3 = 7 → チェックディジットは "7"
func NewCheckDigitByMod10W31(seed seed) checkDigit {
	num := uint64(seed)

	var digits []uint64
	for num > 0 {
		digits = append([]uint64{num % 10}, digits...)
		num /= 10
	}

	var sum uint64
	weights := []uint64{3, 1}
	weightIndex := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * weights[weightIndex]
		weightIndex = (weightIndex + 1) % len(weights)
	}

	mod := sum % 10
	var value uint64
	if mod != 0 {
		value = 10 - mod
	}

	return checkDigit(strconv.FormatUint(value, 10))
}

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

// NewCheckDigitByMod11W10 は、モジュラス11 ウェイト1~0 (M11W1~0) を使用してチェックディジットを計算する。
// 計算の流れは以下の通り:
// 1. データの末尾の桁から順に、1, 2, 3,... と増加するウェイトを掛けて総和を求める。
// 2. 総和を 11 で割った余りを求める。
// 3. 11 から余りを引いた値がチェックディジットとなる。
//   - 余りが 0 または 1 の場合はチェックディジットも 0 になる。
//
// 例: シードが "12345678" の場合
// (8×1)+(7×2)+(6×3)+(5×4)+(4×5)+(3×6)+(2×7)+(1×8) = 8+14+18+20+20+18+14+8 = 120
// 120 % 11 = 10
// 11 - 10 = 1 → チェックディジットは "1"
func NewCheckDigitByMod11W10(seed seed) checkDigit {
	num := uint64(seed)

	var digits []uint64
	for num > 0 {
		digits = append([]uint64{num % 10}, digits...)
		num /= 10
	}

	var sum uint64
	weight := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * uint64(weight)
		weight++
	}

	mod := sum % 11
	var value uint64
	if mod != 0 && mod != 1 {
		value = 11 - mod
	}

	return checkDigit(strconv.FormatUint(value, 10))
}

// NewCheckDigitBySevenCheck は、セブンチェック（7DR及び7DSR）を使用してチェックディジットを計算する。
// 計算の流れは以下の通り:
// 1. データを7で割り、余りを求める。（7DR）
// 2. 7から余りを引いた値がチェックディジットとなる。（7DSR）
//   - 余りが 0 の場合はチェックディジットも 0 になる。
//
// 例: シードが "12345" の場合
// 12345 / 7 = 1763 余り 4 → 7DR: チェックディジットは "4"
// 7 - 4 = 3 → 7DSR: チェックディジットは "3"
func NewCheckDigitBySevenCheck(seed seed) checkDigit {
	num := uint64(seed)

	mod := num % 7
	var value uint64
	if mod != 0 {
		value = 7 - mod
	}

	return checkDigit(strconv.FormatUint(value, 10))
}

// NewCheckDigitByNineCheck は、ナインチェック（9DR及び9DSR）を使用してチェックディジットを計算する。
// 計算の流れは以下の通り:
// 1. データを9で割り、余りを求める。（9DR）
// 2. 9から余りを引いた値がチェックディジットとなる。（9DSR）
//   - 余りが 0 の場合はチェックディジットも 0 になる。
//
// 例: シードが "12345" の場合
// 12345 / 9 = 1371 余り 6 → 9DR: チェックディジットは "6"
// 9 - 6 = 3 → 9DSR: チェックディジットは "3"
func NewCheckDigitByNineCheck(seed seed) checkDigit {
	num := uint64(seed)

	mod := num % 9
	var value uint64
	if mod != 0 {
		value = 9 - mod
	}

	return checkDigit(strconv.FormatUint(value, 10))
}
