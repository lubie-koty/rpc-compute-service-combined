package types

type Server interface {
	Serve() error
}

type MathService interface {
	RootMeanSquare(firstValue float64, secondValue float64) (float64, error)
	GeometricMean(firstValue float64, secondValue float64) (float64, error)
	BodyMassIndex(weight float64, height float64) (float64, error)
	PowerLevelDiff(firstValue float64, secondValue float64) (float64, error)
	PercentageValueChange(firstValue float64, secondValue float64) (float64, error)
}

type SimpleServiceClient interface {
	Add(a, b float64) (float64, error)
	Sub(a, b float64) (float64, error)
	Mul(a, b float64) (float64, error)
	Div(a, b float64) (float64, error)
}

type ComplexServiceClient interface {
	Sqrt(num float64) (float64, error)
	Abs(num float64) (float64, error)
	Power(base float64, exponent float64) (float64, error)
	Log(num float64, base float64) (float64, error)
	Round(num float64, precision float64) (float64, error)
}
