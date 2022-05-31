package payment

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//> model tipe data
type Payment struct {
	PaymentId   int       `gorm : "primaryKey" json:"id"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	PaymentBy   string    `json:"payment"`
	UpdatedAt   time.Time `json:"updated"`
}

//> fungsi get (read) koreksi
func GetAllPayment(c echo.Context) error {
	var product []Payment
	var error error
	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

//> fungsi post(create) koreksi
func CreateProduct(c echo.Context) error {
	var input []Payment
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, input)
}

//> fungsi delete msh butuh koreksi dibagian return
func DeleteProduct(c echo.Context) error {
	var product map[int]*Payment
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}

//>fungsi put(update) msh butuh koreksi
func UpdateProduct(c echo.Context) error {
	var product map[int]*Payment
	n := new(Payment)
	if err := c.Bind(n); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].Description = n.Description
	return c.JSON(http.StatusOK, product[id])
}
