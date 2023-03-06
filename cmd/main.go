package main

import (
	"fmt"
	"hexagonal/internal/adaptors/core/arithmetic"
	"hexagonal/internal/ports"
)

func main() {
	// ports
	var core ports.ArithmeticPort
	core = arithmetic.NewAdapter()

	result, err := core.Addition(1, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
