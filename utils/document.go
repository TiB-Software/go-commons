package utils

import (
	"strings"
	"unicode"
)

func NormalizeDocument(document string) string {
	var builder strings.Builder
	for _, char := range document {
		if unicode.IsDigit(char) {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

func IsValidCPF(document string) bool {
	digits := NormalizeDocument(document)
	if len(digits) != 11 || allDigitsEqual(digits) {
		return false
	}

	firstDigit := cpfCheckDigit(digits[:9], 10)
	secondDigit := cpfCheckDigit(digits[:9]+string(firstDigit), 11)

	return digits[9] == firstDigit && digits[10] == secondDigit
}

func IsValidCNPJ(document string) bool {
	digits := NormalizeDocument(document)
	if len(digits) != 14 || allDigitsEqual(digits) {
		return false
	}

	firstDigit := cnpjCheckDigit(digits[:12], []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	secondDigit := cnpjCheckDigit(digits[:12]+string(firstDigit), []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})

	return digits[12] == firstDigit && digits[13] == secondDigit
}

func cpfCheckDigit(base string, weight int) byte {
	sum := 0
	for _, digit := range base {
		sum += int(digit-'0') * weight
		weight--
	}

	rest := (sum * 10) % 11
	if rest == 10 {
		rest = 0
	}

	return byte(rest) + '0'
}

func cnpjCheckDigit(base string, weights []int) byte {
	sum := 0
	for i, digit := range base {
		sum += int(digit-'0') * weights[i]
	}

	rest := sum % 11
	if rest < 2 {
		return '0'
	}

	return byte(11-rest) + '0'
}

func allDigitsEqual(digits string) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i] != digits[0] {
			return false
		}
	}

	return true
}
