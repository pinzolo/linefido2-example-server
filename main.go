package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pinzolo/linefido2"
	"net/http"
)

const (
	RpId     = "localhost"
	RpOrigin = "http://localhost:4000"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		e.Logger.Error(err)
		res := ErrorResponse{Message: err.Error()}
		c.JSON(http.StatusInternalServerError, res)
	}
	e.POST("/registration", handleRegistration)
	e.POST("/attestation", handleAttestation)
	e.POST("/authentication", handleAuthentication)
	e.POST("/assertion", handleAssertion)
	e.File("/script.js", "./assets/script.js")
	e.File("/registration", "./assets/reg.html")
	e.File("/authentication", "./assets/auth.html")

	client = linefido2.NewClient(
		linefido2.WithRpId(RpId),
		linefido2.WithBaseUrl("http://localhost:8081"))

	e.Logger.Fatal(e.Start(":4000"))
}
