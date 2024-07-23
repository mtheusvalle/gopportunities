package main

import (
	"github.com/mtheusvalle/gopportunities/config"
	"github.com/mtheusvalle/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Error("config initialization error: ", err)
		return
	}

	// Initialize router
	router.Initialize()
}
