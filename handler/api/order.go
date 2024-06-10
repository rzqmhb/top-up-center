package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rzqmhb/top-up-center/service"
)

type OrderAPI interface {
	GetUserOrder(c *gin.Context)
}

type orderAPI struct {
	orderService service.OrderService
	userService service.UserService
}

func NewOrderAPI(orderService service.OrderService, userService service.UserService) OrderAPI {
	return &orderAPI{
		orderService: orderService,
		userService: userService	,
	}
}

func (o *orderAPI) GetUserOrder(c *gin.Context) {
	username, _ := c.Get("username")
	user, err := o.userService.GetByUsername(username.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error", "cause":err.Error()})
		return
	}

	orders, err := o.orderService.GetAllUserOrder(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":"internal server error", "cause":err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}