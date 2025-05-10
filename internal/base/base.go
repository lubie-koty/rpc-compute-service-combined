package base

import "github.com/lubie-koty/rpc-compute-service-combined/internal/core/types"

type CombinedMathService struct {
	simpleService  types.SimpleServiceClient
	complexService types.ComplexServiceClient
}

func NewCombinedMathService(simpleService types.SimpleServiceClient, complexService types.ComplexServiceClient) *CombinedMathService {
	return &CombinedMathService{
		simpleService:  simpleService,
		complexService: complexService,
	}
}

func (s *CombinedMathService) RootMeanSquare(firstValue float64, secondValue float64) float64 {
	xPower := s.complexService.Power(firstValue, 2)
	yPower := s.complexService.Power(secondValue, 2)
	sum := s.simpleService.Add(xPower, yPower)
	avg := s.simpleService.Div(sum, 2)
	return s.complexService.Sqrt(avg)
}

func (s *CombinedMathService) GeometricMean(firstValue float64, secondValue float64) float64 {
	product := s.simpleService.Mul(firstValue, secondValue)
	return s.complexService.Sqrt(product)
}

func (s *CombinedMathService) BodyMassIndex(weight float64, height float64) float64 {
	heightPower := s.complexService.Power(height, 2)
	return s.simpleService.Div(weight, heightPower)
}

func (s *CombinedMathService) PowerLevelDiff(firstValue float64, secondValue float64) float64 {
	ratio := s.simpleService.Div(firstValue, secondValue)
	logValue := s.complexService.Log(ratio, 10)
	return s.simpleService.Mul(10, logValue)
}

func (s *CombinedMathService) PercentageValueChange(firstValue float64, secondValue float64) float64 {
	delta := s.simpleService.Sub(secondValue, firstValue)
	div := s.simpleService.Div(delta, s.complexService.Abs(firstValue))
	percentage := s.simpleService.Mul(div, 100)
	return s.complexService.Round(percentage, 2)
}
