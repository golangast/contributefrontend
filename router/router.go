
package router

import (
	"github.com/labstack/echo/v4"
	"contribute/handler/get/home"

"contribute/handler/get/new"
//#import 
)

//Routes is for routing
func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)


e.GET("/new", new.New)
//#routes
}
