package inertiaMiddleware

import (
	"io"
	"log"
	"text/template"

	"github.com/labstack/echo/v4"
)

type InertiaInfo struct {
	pageTemplate *template.Template
}

func (inertiaInfo *InertiaInfo) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	renderData := map[string]interface{}{}

	return inertiaInfo.pageTemplate.Execute(w, renderData)
}

func RegisterInertiaAdapter(echoInstance *echo.Echo) {
	var inertiaInfo *InertiaInfo = new(InertiaInfo)

	template, err := template.New("app.html").ParseFiles("app.html")
	if err != nil {
		log.Fatal("Failed to load the template", err)
	}

	inertiaInfo.pageTemplate = template
	echoInstance.Renderer = inertiaInfo
}

func InertiaMiddleware(e *echo.Echo) echo.MiddlewareFunc {

	RegisterInertiaAdapter(e)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
