package main

import (
	"flag"
	"medtracker/medtracker/internal/app"
)

func main() {
	configFile := flag.String("config", "./configs/.env", "Path to config file")
	flag.Parse()

	app.Run(*configFile)
}
