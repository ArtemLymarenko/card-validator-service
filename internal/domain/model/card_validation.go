package model

import (
	"time"
)

func isCardDigitValid(digit int) bool {
	const (
		MinNatural = 0
		MaxNatural = 9
	)
	return digit >= MinNatural && digit <= MaxNatural
}

func symbolToInt(symbol byte) int {
	return int(symbol - '0')
}

func (card *Card) isLuhnValid() bool {
	const LuhnBase = 10

	var oddSum, evenSum int
	digits := []byte(card.Number)

	for index, char := range digits {
		digit := symbolToInt(char)
		if !isCardDigitValid(digit) {
			return false
		}

		if index%2 == 0 {
			multiplied := digit * 2
			splitVal := (multiplied % LuhnBase) + (multiplied / LuhnBase)
			evenSum += splitVal
		} else {
			oddSum += digit
		}
	}

	return (evenSum+oddSum)%LuhnBase == 0
}

func (card *Card) isExpired() bool {
	now := time.Now()
	currentMonth := CardMonth(now.Month())
	currentYear := CardYear(now.Year())

	return card.ExpYear < currentYear && card.ExpMonth < currentMonth
}
