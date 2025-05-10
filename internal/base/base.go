package base

import "math"

type ComplexMathService struct{}

func NewComplexMathService() *ComplexMathService {
	return &ComplexMathService{}
}

func (s *ComplexMathService) Sqrt(num float64) float64 {
	return math.Sqrt(num)
}

func (s *ComplexMathService) Abs(num float64) float64 {
	return math.Abs(num)
}

func (s *ComplexMathService) Power(base float64, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func (s *ComplexMathService) Log(num float64, base float64) float64 {
	return math.Log(num) / math.Log(base)
}

func (s *ComplexMathService) Round(num float64, precision int64) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(num*ratio) / ratio
}
