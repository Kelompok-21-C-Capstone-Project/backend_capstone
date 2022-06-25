package transaction

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//> model tipe data
type Transaction struct {
	TransactionId     int       `gorm : "primaryKey" json:"id"`
	UserId            int       `json:"id"`
	PaymentId         int       `json:"id"`
	ProductId         int       `json:"id"`
	Description       string    `json:"description"`
	TransactionDate   time.Time `json:"date"`
	TransactionDetail string    `json:"transaction detail"`
	UpdatedAt         time.Time `json:"updated"`
}

//> fungsi get (read) koreksi
func GetAllTransaction(c echo.Context) error {
	var transaction []Transaction
	var error error
	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}
	return c.JSON(http.StatusOK, transaction)
}

//> fungsi post(create) koreksi
func CreateTransaction(c echo.Context) error {
	var input []Transaction
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, input)
}

//> fungsi delete msh butuh koreksi dibagian return
func DeleteTransaction(c echo.Context) error {
	var transaction map[int]*Transaction
	id, _ := strconv.Atoi(c.Param("id"))
	delete(transaction, id)
	return c.NoContent(http.StatusNoContent)
}

//>fungsi put(update) msh butuh koreksi
func UpdateTransaction(c echo.Context) error {
	var transaction map[int]*Transaction
	n := new(Transaction)
	if err := c.Bind(n); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	transaction[id].Description = n.Description
	transaction[id].TransactionDate = n.TransactionDate
	transaction[id].TransactionDetail = n.TransactionDetail
	return c.JSON(http.StatusOK, transaction[id])
}
