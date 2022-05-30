package main

import (
	"github.com/izaakdale/goBank/app"
	"github.com/izaakdale/utils-go/logger"
)

func main() {
	logger.Info("Starting Application")
	app.Start()
}
