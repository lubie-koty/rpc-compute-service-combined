package types

type Server interface {
	Serve() error
}

type MathService interface {
	RootMeanSquare(numbers []float64) (float64, error)
	GeometricMean(numbers []float64) (float64, error)
	BodyMassIndex(weight float64, height float64) (float64, error)
	PowerLevelDiff(firstValue float64, secondValue float64) (float64, error)
	PercentageValueChange(firstValue float64, secondValue float64) (float64, error)
}

type SimpleServiceClient interface {
	Add(numbers []float64) (float64, error)
	Sub(numbers []float64) (float64, error)
	Mul(numbers []float64) (float64, error)
	Div(numbers []float64) (float64, error)
}

type ComplexServiceClient interface {
	Sqrt(num float64) (float64, error)
	Abs(num float64) (float64, error)
	Power(base float64, exponent float64) (float64, error)
	Log(num float64, base float64) (float64, error)
	Round(num float64, precision float64) (float64, error)
}
