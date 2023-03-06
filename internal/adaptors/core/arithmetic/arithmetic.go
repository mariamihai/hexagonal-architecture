package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith Adapter) Addition(a int, b int) (int, error) {
	return a + b, nil
}

func (arith Adapter) Subtraction(a int, b int) (int, error) {
	return a - b, nil
}

func (arith Adapter) Multiplication(a int, b int) (int, error) {
	return a * b, nil
}

func (arith Adapter) Division(a int, b int) (int, error) {
	return a / b, nil
}
