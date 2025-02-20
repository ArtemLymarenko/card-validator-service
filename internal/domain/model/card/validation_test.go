package card

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsCardDigitValid(t *testing.T) {
	tests := []struct {
		digit   int
		isValid bool
	}{
		{digit: 0, isValid: true},
		{digit: 5, isValid: true},
		{digit: 9, isValid: true},
		{digit: -1, isValid: false},
		{digit: 10, isValid: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.digit), func(t *testing.T) {
			assert.Equal(t, test.isValid, isCardDigitValid(test.digit))
		})
	}
}

func TestSymbolToInt(t *testing.T) {
	tests := []struct {
		symbol byte
		result int
	}{
		{symbol: '0', result: 0},
		{symbol: '5', result: 5},
		{symbol: '9', result: 9},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.symbol), func(t *testing.T) {
			assert.Equal(t, test.result, symbolToInt(test.symbol))
		})
	}
}

func TestIsValidNumberLuhn(t *testing.T) {
	tests := []struct {
		card    Card
		isValid bool
	}{
		{card: Card{Number: "4111111111111111"}, isValid: true},
		{card: Card{Number: "371449635398431"}, isValid: true},
		{card: Card{Number: "30569309025904"}, isValid: true},
		{card: Card{Number: "5555555555554444"}, isValid: true},
		{card: Card{Number: "4111111111111112"}, isValid: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.card.Number), func(t *testing.T) {
			assert.Equal(t, test.isValid, test.card.IsValidNumberLuhn())
		})
	}
}

func TestIsValidYear(t *testing.T) {
	tests := []struct {
		card    Card
		isValid bool
	}{
		{card: Card{ExpYear: Year(time.Now().Year())}, isValid: true},
		{card: Card{ExpYear: Year(time.Now().Year() + 1)}, isValid: true},
		{card: Card{ExpYear: Year(2099)}, isValid: true},
		{card: Card{ExpYear: Year(2100)}, isValid: false},
		{card: Card{ExpYear: Year(time.Now().Year() - 1)}, isValid: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.card.ExpYear), func(t *testing.T) {
			assert.Equal(t, test.isValid, test.card.IsValidYear())
		})
	}
}

func TestIsValidMonth(t *testing.T) {
	tests := []struct {
		card    Card
		isValid bool
	}{
		{card: Card{ExpMonth: Month(time.Now().Month())}, isValid: true},
		{card: Card{ExpMonth: Month(5)}, isValid: true},
		{card: Card{ExpMonth: Month(12)}, isValid: true},
		{card: Card{ExpMonth: Month(13)}, isValid: false},
		{card: Card{ExpMonth: Month(0)}, isValid: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.card.ExpMonth), func(t *testing.T) {
			assert.Equal(t, test.isValid, test.card.IsValidMonth())
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		card    Card
		isValid bool
		err     error
	}{
		{
			card:    Card{Number: "4111111111111111", ExpYear: Year(time.Now().Year() + 1), ExpMonth: Month(time.Now().Month() + 1)},
			isValid: true,
			err:     nil,
		},
		{
			card:    Card{Number: "4111111111111112", ExpYear: Year(time.Now().Year() + 1), ExpMonth: Month(5)},
			isValid: false,
			err:     ErrInvalidLuhnCheck,
		},
		{
			card:    Card{Number: "4111111111111111", ExpYear: Year(time.Now().Year() - 1), ExpMonth: Month(5)},
			isValid: false,
			err:     ErrInvalidCardYear,
		},
		{
			card:    Card{Number: "4111111111111111", ExpYear: Year(time.Now().Year()), ExpMonth: Month(time.Now().Month() - 1)},
			isValid: false,
			err:     ErrInvalidCardMonth,
		},
		{
			card:    Card{Number: "371449635398431", ExpYear: Year(time.Now().Year()), ExpMonth: Month(time.Now().Month())},
			isValid: true,
			err:     nil,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.card.Number), func(t *testing.T) {
			valid, err := test.card.Validate()
			assert.Equal(t, test.isValid, valid)
			assert.Equal(t, test.err, err)
		})
	}
}
