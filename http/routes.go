package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/sukrati192/fileserver/docs"
	"github.com/sukrati192/fileserver/handlers"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/file/:fileID", handlers.GetFile)
	e.POST("/upload", handlers.UploadFile)
}
