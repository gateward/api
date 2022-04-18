package main

import (
	"api/confs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := confs.ConnectDB("127.0.0.1", "postgres", "postgres", "postgres", "5432")
	if err != nil {
		panic("cannot connect to postgresql")
	}
	confs.MigrateUsersOrgs(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
