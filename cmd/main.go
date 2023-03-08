package main

import (
	"fmt"
	"hexagonal/internal/adaptors/app/api"
	"hexagonal/internal/adaptors/core/arithmetic"
	"hexagonal/internal/adaptors/framework/driven/db"
	"hexagonal/internal/ports"
	"log"
)

func main() {
	// core
	var core ports.ArithmeticPort
	core = arithmetic.NewAdapter()

	// framework layer
	// DB
	var dbAdapter ports.DBPort
	dbAdapter, err := db.NewAdapter("driverName", "dataSourceName")
	if err != nil {
		log.Fatal(err)
	}

	// application layer
	var app ports.APIPort
	app = api.NewAdapter(core, dbAdapter)

	result, err := app.GetAddition(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
