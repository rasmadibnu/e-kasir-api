package main

import (
	"flag"
	"kasir-cepat-api/config"
	"kasir-cepat-api/routes"
)

func main() {
	config.LoadEnv()
	db := config.InitialDB()
	// config.Migration(db)

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		config.InitCommands(db)
	} else {
		routes.WebRouter(db)
	}
}
