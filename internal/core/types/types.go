package types

type Server interface {
	Serve() error
}

type MathService interface {
	RootMeanSquare(firstValue float64, secondValue float64) float64
	GeometricMean(firstValue float64, secondValue float64) float64
	BodyMassIndex(weight float64, height float64) float64
	PowerLevelDiff(firstValue float64, secondValue float64) float64
	PercentageValueChange(firstValue float64, secondValue float64) float64
}

type SimpleServiceClient interface {
	Add(a, b float64) float64
	Sub(a, b float64) float64
	Mul(a, b float64) float64
	Div(a, b float64) float64
}

type ComplexServiceClient interface {
	Sqrt(num float64) float64
	Abs(num float64) float64
	Power(base float64, exponent float64) float64
	Log(num float64, base float64) float64
	Round(num float64, precision int64) float64
}
