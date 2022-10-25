package product_domain

import (
	"products/utils/error_utils"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name" valid:"required~name is required"`
	Price     float32   `json:"price" valid:"required~price is required"`
	Stock     int8      `json:"stock" valid:"required~stock is required"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error_utils.MessageErr {
	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	return nil
}

func (p *Product) GetProductIdParam(c *gin.Context) (int64, error_utils.MessageErr) {
	idParam := c.Param("productId")
	movieId, err := strconv.Atoi(idParam)
	if err != nil {
		return int64(0), error_utils.NewBadRequest("invalid product id params")
	}

	return int64(movieId), nil
}
