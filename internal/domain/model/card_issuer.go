package model

type IssuerName string

type IIN int

type CardIssuer struct {
	Name         IssuerName
	IINs         []IIN
	DigitsAmount []int
}
