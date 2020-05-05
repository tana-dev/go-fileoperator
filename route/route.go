package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rakyll/statik/fs"
	"github.com/tana-dev/go-filesplitter/api"
	_ "github.com/tana-dev/go-filesplitter/statik"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// statik
	statikFS, _ := fs.New()
	h := http.FileServer(statikFS)
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", h)))

	// Routes
	apiGroup := e.Group("/api/v1")
	{
		apiGroup.POST("/filesplit", api.PostFilesplit)
	}

	return e
}
