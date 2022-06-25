package product

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//> model tipe data
type Product struct {
	ProductId   int       `gorm : "primaryKey" json:"id"`
	ProductName string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated"`
}

//> fungsi get (read) koreksi
func GetAllProduct(c echo.Context) error {
	var product []Product
	var error error
	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

//> fungsi post(create) koreksi
func CreateProduct(c echo.Context) error {
	var input []Product
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, input)
}

//> fungsi delete msh butuh koreksi dibagian return
func DeleteProduct(c echo.Context) error {
	var product map[int]*Product
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}

//>fungsi put(update) msh butuh koreksi
func UpdateProduct(c echo.Context) error {
	var product map[int]*Product
	n := new(Product)
	if err := c.Bind(n); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].ProductName = n.ProductName
	return c.JSON(http.StatusOK, product[id])
}
