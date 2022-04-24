package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	zxcvbn "github.com/nbutton23/zxcvbn-go"
)

type goTrueError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func CreateUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if password != c.FormValue("password-verif") {
		return echo.NewHTTPError(http.StatusBadRequest, "Password and password verification does not match")
	}

	strength := zxcvbn.PasswordStrength(password, nil)
	if strength.Score < 2 {
		return echo.NewHTTPError(http.StatusBadRequest, "Password is too weak")
	}

	body, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to create user, please contact an administrator")
	}

	resp, err := http.Post("http://127.0.0.1:9999/signup", "application/json", bytes.NewBuffer(body))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var body goTrueError
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&body)
		if err != nil {
			c.Logger().Error(err)
		}
		return echo.NewHTTPError(resp.StatusCode, body.Message)
	}

	c.Logger().Info("User " + email + " created")

	return c.String(http.StatusOK, "ok")
}
