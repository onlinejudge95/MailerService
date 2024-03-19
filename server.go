package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func StartGin(conf Config) {
	switch strings.ToUpper(conf.Environment) {
	case "DEBUG":
		gin.SetMode(gin.DebugMode)
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)
	case "TEST":
		gin.SetMode(gin.TestMode)
	default:
		panic("Unidentified mode for running the server")
	}

	gin.DisableConsoleColor()

	router := gin.Default()
	router.POST("/contact-us/", ContactUsRouter)
	router.Run(conf.BindAddress)
}
