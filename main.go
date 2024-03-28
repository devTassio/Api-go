package main

import (
	"fmt"

	"github.com/devTassio/Api-go/config"
	"github.com/devTassio/Api-go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
