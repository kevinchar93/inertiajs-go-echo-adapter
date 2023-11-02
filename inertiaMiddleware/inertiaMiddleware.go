package inertiaMiddleware

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const (
	HEADER_X_INERTIA                   = "X-Inertia"
	HEADER_X_INERTIA_VERSION           = "X-Inertia-Version"
	HEADER_X_INERTIA_LOCATION          = "X-Inertia-Location"
	HEADER_X_INERTIA_PARTIAL_DATA      = "X-Inertia-Partial-Data"
	HEADER_X_INERTIA_PARTIAL_COMPONENT = "X-Inertia-Partial-Component"
)

func hasHeader(req *http.Request, header string) bool {
	return req.Header.Get(header) != ""
}

type PageTemplateAssets struct {
	jsFiles  []template.HTMLAttr
	cssFiles []template.HTMLAttr
}

func NewPageTemplateAssets() (p *PageTemplateAssets) {
	p = new(PageTemplateAssets)

	if os.Getenv("BUILD_ENV") == "development" {
		p.jsFiles = append(p.jsFiles,
			template.HTMLAttr("http://localhost:5173/@vite/client"),
			template.HTMLAttr("http://localhost:5173/src/main.jsx"))

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

	if hasHeader(req, HEADER_X_INERTIA) {
		ctx.Response().Header().Set(HEADER_X_INERTIA, "true")
		return ctx.JSON(http.StatusOK, page)
	}

	pageJson, err := json.Marshal(page)
	if err != nil {
		log.Fatal("Failed to encode page object to JSON", err)
	}

	renderData := map[string]interface{}{
		"pageObject":    template.HTMLAttr(string(pageJson)),
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
