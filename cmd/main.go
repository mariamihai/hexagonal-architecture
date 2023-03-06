package main

import (
	"fmt"
	"hexagonal/internal/adaptors/app/api"
	"hexagonal/internal/adaptors/core/arithmetic"
	"hexagonal/internal/ports"
)

func main() {
	// ports
	var core ports.ArithmeticPort
	core = arithmetic.NewAdapter()

	var app ports.APIPort
	app = api.NewAdapter(core)

	result, err := app.GetAddition(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
