package inertiaMiddleware

import (
	"encoding/json"
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
	assetVersion       string
}

func RegisterInertiaAdapter(echoInstance *echo.Echo, assetVersion string) {
	var inertiaInfo *InertiaInfo = new(InertiaInfo)

	inertiaInfo.assetVersion = assetVersion
	inertiaInfo.pageTemplateAssets = NewPageTemplateAssets()

	template, err := template.New("app.html").ParseFiles("app.html")
	if err != nil {
		log.Fatal("Failed to load the template", err)
	}

	inertiaInfo.pageTemplate = template
	echoInstance.Renderer = inertiaInfo
}

func (inertiaInfo *InertiaInfo) Render(w io.Writer, name string, props interface{}, ctx echo.Context) error {

	req := ctx.Request()

	page := map[string]interface{}{
		"component": name,
		"props":     props.(map[string]interface{}),
		"url":       req.URL.String(),
		"version":   inertiaInfo.assetVersion,
	}

	pageJson, err := json.Marshal(page)
	if err != nil {
		log.Fatal("Failed to encode page object to JSON", err)
	}

	renderData := map[string]interface{}{
		"pageObject":    string(pageJson),
		"jsFiles":       inertiaInfo.pageTemplateAssets.jsFiles,
		"cssFiles":      inertiaInfo.pageTemplateAssets.cssFiles,
		"isDevelopment": os.Getenv("BUILD_ENV") == "development",
	}

	return inertiaInfo.pageTemplate.Execute(w, renderData)
}

func InertiaMiddleware(e *echo.Echo, assetVersion string) echo.MiddlewareFunc {

	RegisterInertiaAdapter(e, assetVersion)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}
