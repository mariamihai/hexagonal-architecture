package api

import "hexagonal/internal/ports"

type Adapter struct {
	arithPort ports.ArithmeticPort
}

func NewAdapter(arithPort ports.ArithmeticPort) *Adapter {
	return &Adapter{
		arithPort: arithPort,
	}
}

func (adapt Adapter) GetAddition(a, b int) (int, error) {
	answer, err := adapt.arithPort.Addition(a, b)

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapt Adapter) GetSubtraction(a, b int) (int, error) {
	answer, err := adapt.arithPort.Subtraction(a, b)

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapt Adapter) GetMultiplication(a, b int) (int, error) {
	answer, err := adapt.arithPort.Multiplication(a, b)

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (adapt Adapter) GetDivision(a, b int) (int, error) {
	answer, err := adapt.arithPort.Division(a, b)

	if err != nil {
		return 0, err
	}

	return answer, nil
}
