package main

import (
	"log"
	"net/http"

	"codeforge/config"
	"codeforge/server"

	"github.com/gin-gonic/gin"
)

func init() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error in loading config file: %v", err)
	}
}

func main() {
	routes := gin.New()

	log.Printf("server started on port: %v", config.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+config.Cfg.ServerPort, server.NewServerImpl(routes)))
}
