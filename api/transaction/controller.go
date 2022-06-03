package transaction

import (
	"backend_capstone/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//> fungsi get (read) koreksi
func GetAllTransaction(c echo.Context) error {
	var transaction []models.Transaction
	var error error
	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}
	return c.JSON(http.StatusOK, transaction)
}

//> fungsi post(create) koreksi
func CreateTransaction(c echo.Context) error {
	var input []models.Transaction
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, input)
}

//> fungsi delete msh butuh koreksi dibagian return
func DeleteTransaction(c echo.Context) error {
	var transaction map[int]*models.Transaction
	id, _ := strconv.Atoi(c.Param("id"))
	delete(transaction, id)
	return c.NoContent(http.StatusNoContent)
}

//>fungsi put(update) msh butuh koreksi
func UpdateTransaction(c echo.Context) error {
	var transaction map[int]*models.Transaction
	n := new(models.Transaction)
	if err := c.Bind(n); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	transaction[id].Description = n.Description
	transaction[id].TransactionDate = n.TransactionDate
	transaction[id].TransactionDetail = n.TransactionDetail
	return c.JSON(http.StatusOK, transaction[id])
}
