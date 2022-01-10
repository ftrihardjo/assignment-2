package main

import (
	"assignment-2/db"
	"assignment-2/router"
)

const PORT = ":8080"

func init() {
	db.StartDB("localhost", 5432, "postgres", "postgres", "orders_by")
}

func main() {
	r := router.StartApp()
	r.Run(PORT)
	db.GetDB().Close()
}
