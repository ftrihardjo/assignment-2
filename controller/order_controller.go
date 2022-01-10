package controller

import (
	"assignment-2/db"
	"assignment-2/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := db.GetDB()
	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		panic(err.Error())
	}
	statement := "INSERT INTO orders (customer_name,ordered_at) VALUES ($1, $2) RETURNING order_id"
	orderTime, err := time.Parse("2006-01-02T15:04:05-0700", order.OrderedAt)
	if err != nil {
		panic(err.Error())
	}
	row := db.QueryRow(statement, order.CustomerName, orderTime)
	var orderID uint
	row.Scan(&orderID)
	statement = "INSERT INTO items (item_code,description,quantity,order_id) VALUES ($1, $2, $3, $4)"
	for _, item := range order.Items {
		db.QueryRow(statement, item.ItemCode, item.Description, item.Quantity, orderID)
	}
}

func GetOrders(c *gin.Context) {
	db := db.GetDB()
	orderStatement := "SELECT * FROM orders"
	orders, err := db.Query(orderStatement)
	defer orders.Close()
	if err != nil {
		panic(err.Error())
	}
	var result []models.Order
	for orders.Next() {
		order := models.Order{}
		orders.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt)
		itemStatement := "SELECT item_id,item_code,description,quantity FROM items WHERE order_id=$1"
		items, err := db.Query(itemStatement, order.OrderId)
		if err != nil {
			panic(err.Error())
		}
		for items.Next() {
			item := models.Item{}
			items.Scan(&item.LineItemId, &item.ItemCode, &item.Description, &item.Quantity)
			order.Items = append(order.Items, item)
		}
		items.Close()
		result = append(result, order)
	}
	c.JSON(http.StatusOK, result)
}

func UpdateOrder(c *gin.Context) {
	db := db.GetDB()
	order := models.Order{}
	if err := c.ShouldBindJSON(&order); err != nil {
		panic(err.Error())
	}
	orderID, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		panic(err.Error())
	}
	order.OrderId = uint(orderID)
	statement := "UPDATE orders SET customer_name=$1, ordered_at=$2 WHERE order_id=$3"
	db.QueryRow(statement, order.CustomerName, order.OrderedAt, order.OrderId)
	for _, item := range order.Items {
		statement = "UPDATE items SET item_code=$1, description=$2, quantity=$3 WHERE item_id=$4"
		db.QueryRow(statement, item.ItemCode, item.Description, item.Quantity, item.LineItemId)
	}
}

func DeleteOrder(c *gin.Context) {
	db := db.GetDB()
	orderID, err := strconv.Atoi(c.Param("orderId"))
	if err != nil {
		panic(err.Error())
	}
	statement := "DELETE FROM orders WHERE order_id=$1"
	db.QueryRow(statement, orderID)
	statement = "DELETE FROM items WHERE order_id=$1"
	db.QueryRow(statement, orderID)
}
