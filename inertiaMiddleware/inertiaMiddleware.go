package inertiaMiddleware

import (
	"log"
	"text/template"

	"github.com/labstack/echo/v4"
)

type InertiaInfo struct {
	pageTemplate *template.Template
}

func RegisterInertiaAdapter(echoInstance *echo.Echo) {
	var inertiaInfo *InertiaInfo = new(InertiaInfo)

	template, err := template.New("app.html").ParseFiles("app.html")
	if err != nil {
		log.Fatal("Failed to load the template", err)
	}

	inertiaInfo.pageTemplate = template
}

func InertiaMiddleware(e *echo.Echo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
