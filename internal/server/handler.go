package server

import (
	"Search/internal/search"
	"log"
	"net/http"

	"github.com/labstack/echo"
)


func registerRoutes(e *echo.Echo) {
	e.GET("/search", searchHandler)
}

func searchHandler(c echo.Context) error {
	query := c.QueryParam("q")
	results, err := search.SearchPruducts(query)
	if err != nil {
		log.Println("Fa")
		c.String(http.StatusInternalServerError, "")
	}

	

	return c.JSON(http.StatusOK, results)
}
