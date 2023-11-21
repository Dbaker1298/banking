package main

import (
	"github.com/Dbaker1298/banking/app"
	"github.com/Dbaker1298/banking/logger"
)

func main() {
	// log.Println("Starting the application...")
	logger.Info("Starting the application...")
	app.Start()
}
