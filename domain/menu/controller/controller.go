package controller

import (
	"github.com/gin-gonic/gin"
	"jummechu/domain/menu"
	"net/http"
)

type MenuController struct {
}

func (m MenuController) GetMenu(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": menu.Menus,
	})
}
