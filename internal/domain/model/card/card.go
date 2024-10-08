package card

type Number string

type Year int

type Month int

type Card struct {
	Number   Number
	ExpYear  Year
	ExpMonth Month
}
