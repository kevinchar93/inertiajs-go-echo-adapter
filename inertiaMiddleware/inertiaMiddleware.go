package inertiaMiddleware

import (
	"io"
	"log"
	"os"
	"text/template"

	"github.com/labstack/echo/v4"
)

type PageTemplateAssets struct {
	jsFiles  []string
	cssFiles []string
}

func NewPageTemplateAssets() (p *PageTemplateAssets) {
	p = new(PageTemplateAssets)

	if os.Getenv("BUILD_ENV") == "development" {
		p.jsFiles = append(p.jsFiles,
			"http://localhost:5173/@vite/client",
			"http://localhost:5173/src/main.jsx")

		return p
	}
	return p
}

type InertiaInfo struct {
	pageTemplate       *template.Template
	pageTemplateAssets *PageTemplateAssets
}

func RegisterInertiaAdapter(echoInstance *echo.Echo) {
	var inertiaInfo *InertiaInfo = new(InertiaInfo)

	inertiaInfo.pageTemplateAssets = NewPageTemplateAssets()

	template, err := template.New("app.html").ParseFiles("app.html")
	if err != nil {
		log.Fatal("Failed to load the template", err)
	}

	inertiaInfo.pageTemplate = template
	echoInstance.Renderer = inertiaInfo
}

func (inertiaInfo *InertiaInfo) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	renderData := map[string]interface{}{
		"jsFiles":       inertiaInfo.pageTemplateAssets.jsFiles,
		"cssFiles":      inertiaInfo.pageTemplateAssets.cssFiles,
		"isDevelopment": os.Getenv("BUILD_ENV") == "development",
	}

	return inertiaInfo.pageTemplate.Execute(w, renderData)
}

func InertiaMiddleware(e *echo.Echo) echo.MiddlewareFunc {

	RegisterInertiaAdapter(e)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
