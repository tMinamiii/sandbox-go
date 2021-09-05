//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package sgomock

type Arithmetic interface {
	Add(x int64, y int64) int64
	Sub(x int64, y int64) int64
}

type ArithmeticImpl struct {
}

func NewArithmetic() Arithmetic {
	return &ArithmeticImpl{}
}

func (a *ArithmeticImpl) Add(x int64, y int64) int64 {
	return x + y
}

func (a *ArithmeticImpl) Sub(x int64, y int64) int64 {
	return x - y
}
