package models

type Order struct {
	OrderId      uint
	CustomerName string
	OrderedAt    string
	Items        []Item
}
