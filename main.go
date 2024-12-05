package main

import (
	"my-go-project/internal/config"
	"my-go-project/router"
)

func main() {
	config.InitDB()
	r := router.SetupRouter()
	r.Run(":8080")
}
