package types

type Server interface {
	Serve() error
}

type MathService interface {
	RootMeanSquare(firstValue float64, secondValue float64) float64
	GeometricMean(firstValue float64, secondValue float64) float64
	BodyMassIndex(firstValue float64, secondValue float64) float64
	PowerLevelDiff(firstValue float64, secondValue float64) float64
	PercentageValueChange(firstValue float64, secondValue float64) float64
}
