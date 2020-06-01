package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "API OK")
	})

	e.GET("/users", func(c echo.Context) error {
		var res struct {
			Status	string `json:"status"`
			Message string `json:"message"`
		}

		res.Status = "Ok"
		res.Message = "Retorna todos los usuarios"

		return c.JSON(http.StatusOK, &res)
	})

	e.GET("/users/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Datos de un usuario")
	})

	e.POST("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Crea un nuevo usuario")
	})

	e.PUT("/users/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Actualizar datos del usuario")
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Eliminar registro de usuario")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

