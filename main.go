package main

import (
	"api/confs"
	"api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	env "github.com/owlint/go-env"
)

func main() {
	e := echo.New()

	host := env.GetMandatoryEnv("MYSQL_HOST")
	user := env.GetMandatoryEnv("MYSQL_USER")
	pass := env.GetMandatoryEnv("MYSQL_PASS")
	dbName := env.GetDefaultEnv("MYSQL_DB", "gateward")
	port := env.GetDefaultEnv("MYSQL_PORT", "3306")

	e.Logger.SetLevel(log.INFO)

	db, err := confs.ConnectDB(host, user, pass, dbName, port)
	if err != nil {
		panic("cannot connect to mysql")
	}
	confs.MigrateUsersOrgs(db)

	// Users routes (GoTrue gateway)
	e.POST("/user", controllers.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
