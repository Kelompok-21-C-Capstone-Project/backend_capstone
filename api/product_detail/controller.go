package productdetail

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//> model tipe data
type Detail struct {
	ProductDetailId int       `gorm : "primaryKey" json:"id"`
	ProductId       int       `json:"id"`
	ProductList     string    `json:"list"`
	Description     string    `json:"description"`
	StockProduct    int       `json:"stock"`
	UpdatedAt       time.Time `json:"updated"`
}

//> fungsi get (read) koreksi
func GetAllDetail(c echo.Context) error {
	var product []Detail
	var error error
	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

//> fungsi post(create) koreksi
func CreateProductDetail(c echo.Context) error {
	var input []Detail
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, input)
}

//> fungsi delete msh butuh koreksi dibagian return
func DeleteProductDetail(c echo.Context) error {
	var product map[int]*Detail
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}

//>fungsi put(update) msh butuh koreksi
func UpdateProductDetail(c echo.Context) error {
	var product map[int]*Detail
	n := new(Detail)
	if err := c.Bind(n); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].StockProduct = n.StockProduct
	product[id].Description = n.Description
	return c.JSON(http.StatusOK, product[id])
}
