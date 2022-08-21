package sgomock

type Calculator interface {
	Calc(x int64, y int64) int64
}

type CalculatorImpl struct {
	ArithmeticService Arithmetic
}

func NewCalculator() Calculator {
	return &CalculatorImpl{
		ArithmeticService: NewArithmetic(),
	}
}

func (c *CalculatorImpl) Calc(x int64, y int64) int64 {
	add := c.ArithmeticService.Add(x, y)
	return c.ArithmeticService.Sub(add, 10)
}
