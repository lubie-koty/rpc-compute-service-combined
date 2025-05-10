package types

type Server interface {
	Serve() error
}

type MathService interface {
	Sqrt(num float64) float64
	Abs(num float64) float64
	Power(base float64, exponent float64) float64
	Log(num float64, base float64) float64
	Round(num float64, precision int64) float64
}
