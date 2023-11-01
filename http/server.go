package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sukrati192/fileserver/configmanager"
	"github.com/sukrati192/fileserver/services"
)

var (
	host           = flag.String("host", "0.0.0.0", "Host ip")
	port           = flag.String("port", "8000", "Host port")
	configFilePath = flag.String("config", "../configmanager/configs/config.json", "config file path")
)

// @title FileServer
// @version 1.0
// @description A Service that facilitates file upload and download using IPFS
// @host localhost:8000
// @BasePath /
func main() {
	httpServer := echo.New()
	AddRoutes(httpServer)

	flag.Parse()
	if err := configmanager.LoadConfiguration(configFilePath); err != nil {
		log.Fatal("Failed initializing configmanager. Error = ", err)
	}

	services.InitIPFS()

	if err := httpServer.Start(fmt.Sprintf("%s:%s", *host, *port)); err != nil {
		log.Fatal("Failed to start server!", err)
	}
}
