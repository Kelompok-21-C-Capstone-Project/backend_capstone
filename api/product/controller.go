package product

import (
	"backend_capstone/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//> fungsi get (read) koreksi
func GetAllProduct(c echo.Context) error {
	var product []models.Product
	var error error
	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

//> fungsi post(create) koreksi
func CreateProduct(c echo.Context) error {
	var input []models.Product
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, input)
}

//> fungsi delete msh butuh koreksi dibagian return
func DeleteProduct(c echo.Context) error {
	var product map[int]*models.Product
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}

//>fungsi put(update) msh butuh koreksi
func UpdateProduct(c echo.Context) error {
	var product map[int]*models.Product
	n := new(models.Product)
	if err := c.Bind(n); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].ProductName = n.ProductName
	return c.JSON(http.StatusOK, product[id])
}
