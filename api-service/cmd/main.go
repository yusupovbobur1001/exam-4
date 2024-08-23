package main

import (
	"api_service/api"
	"api_service/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.Load()
	fmt.Printf("api service run %s ...", cfg.HTTP_PORT)
	router := api.NewRouter(cfg)
	err := router.Run(cfg.HTTP_PORT)
	if err != nil {
		panic(err)
	}
}
