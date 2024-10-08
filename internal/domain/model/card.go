package model

type CardNumber string

type CardYear int

type CardMonth int

type Card struct {
	Number   CardNumber
	ExpYear  CardYear
	ExpMonth CardMonth
}
