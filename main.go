package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/onlinejudge95/MailerService/pkg/mailer"
)

func main() {
	conf, err := mailer.LoadConfig(".")

	if err != nil {
		fmt.Println("Error occured while loading config", err)
	}

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
	router.POST("/contact-us/", mailer.ContactUsRouter)
	router.Run(conf.BindAddress)
}
