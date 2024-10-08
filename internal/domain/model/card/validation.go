package card

import (
	"errors"
	"time"
)

var (
	ErrInvalidCardNumber = errors.New("invalid card number")
	ErrInvalidLuhnCheck  = errors.Join(ErrInvalidCardNumber, errors.New("luhn check has failed"))
	ErrInvalidCardYear   = errors.New("invalid card year")
	ErrInvalidCardMonth  = errors.New("invalid card month")
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

func (c Card) IsValidNumberLuhn() bool {
	const LuhnBase = 10

	var totalSum int
	digits := []byte(c.Number)

	isSecondDigit := false
	for i := len(digits) - 1; i >= 0; i-- {
		digit := symbolToInt(digits[i])
		if !isCardDigitValid(digit) {
			return false
		}

		if isSecondDigit {
			digit *= 2
			digit = digit%LuhnBase + digit/LuhnBase
		}

		totalSum += digit

		isSecondDigit = !isSecondDigit
	}

	return (totalSum)%LuhnBase == 0
}

func (c Card) IsValidYear() bool {
	const (
		MaxValidYear = Year(2099)
	)

	currentYear := Year(time.Now().Year())
	return c.ExpYear >= currentYear && c.ExpYear <= MaxValidYear
}

func (c Card) IsValidMonth() bool {
	const (
		MinValidMonth = Month(1)
		MaxValidMonth = Month(12)
	)

	currentYear := Year(time.Now().Year())
	currentMonth := Month(time.Now().Month())

	if c.ExpYear == currentYear {
		return c.ExpMonth >= currentMonth && c.ExpMonth <= MaxValidMonth
	}

	return c.ExpMonth >= MinValidMonth && c.ExpMonth <= MaxValidMonth
}

func (c Card) Validate() (valid bool, err error) {
	if valid = c.IsValidYear(); !valid {
		return valid, ErrInvalidCardYear
	}

	if valid = c.IsValidMonth(); !valid {
		return valid, ErrInvalidCardMonth
	}

	if valid = c.IsValidNumberLuhn(); !valid {
		return valid, ErrInvalidLuhnCheck
	}

	return valid, nil
}
