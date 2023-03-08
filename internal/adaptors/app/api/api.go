package api

import "hexagonal/internal/ports"

type Adapter struct {
	arithmeticPort ports.ArithmeticPort
	dbPort         ports.DBPort
}

func NewAdapter(arithPort ports.ArithmeticPort, dbPort ports.DBPort) *Adapter {
	return &Adapter{
		arithmeticPort: arithPort,
		dbPort:         dbPort,
	}
}

func (apiAdapter Adapter) GetAddition(a, b int) (int, error) {
	answer, err := apiAdapter.arithmeticPort.Addition(a, b)

	if err != nil {
		return 0, err
	}

	err = apiAdapter.dbPort.AddToHistory(answer, "addition")

	return answer, nil
}

func (apiAdapter Adapter) GetSubtraction(a, b int) (int, error) {
	answer, err := apiAdapter.arithmeticPort.Subtraction(a, b)

	if err != nil {
		return 0, err
	}

	err = apiAdapter.dbPort.AddToHistory(answer, "subtraction")

	return answer, nil
}

func (apiAdapter Adapter) GetMultiplication(a, b int) (int, error) {
	answer, err := apiAdapter.arithmeticPort.Multiplication(a, b)

	if err != nil {
		return 0, err
	}

	err = apiAdapter.dbPort.AddToHistory(answer, "multiplication")

	return answer, nil
}

func (apiAdapter Adapter) GetDivision(a, b int) (int, error) {
	answer, err := apiAdapter.arithmeticPort.Division(a, b)

	if err != nil {
		return 0, err
	}

	err = apiAdapter.dbPort.AddToHistory(answer, "division")

	return answer, nil
}
