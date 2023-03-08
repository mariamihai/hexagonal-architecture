package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arithmeticAdapter Adapter) Addition(a int, b int) (int, error) {
	return a + b, nil
}

func (arithmeticAdapter Adapter) Subtraction(a int, b int) (int, error) {
	return a - b, nil
}

func (arithmeticAdapter Adapter) Multiplication(a int, b int) (int, error) {
	return a * b, nil
}

func (arithmeticAdapter Adapter) Division(a int, b int) (int, error) {
	return a / b, nil
}
