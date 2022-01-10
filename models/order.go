package models

import "time"

type Order struct {
	OrderId      uint
	CustomerName string
	OrderedAt    time.Time `json:"ordered_at`
	Items        []Item
}
