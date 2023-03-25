
package new

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(c echo.Context) error {
	var new = "new"
	return c.Render(http.StatusOK, "new.html", map[string]interface{}{
		"new":new,
	})

}