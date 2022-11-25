package routes

import (
	"jummechu/domain/menu/controller"
	"jummechu/lib"
)

type MenuRoutes struct {
	server         lib.Server
	menuController controller.MenuController
}

func (m MenuRoutes) Setup() {
	api := m.server.Gin.Group("")
	{
		api.GET("/menu", m.menuController.GetMenu)
	}
}

func NewMenuRoutes(s lib.Server) MenuRoutes {
	return MenuRoutes{
		server: s,
	}
}
