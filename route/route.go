package route

import (
	"github.com/tana-dev/go-filesplitter/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	// Set Bundle MiddleWare
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	//	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//		AllowOrigins: []string{"*"},
	//		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	//	}))

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/filesplit", handler.PostFilesplit)
	}

	return e
}
