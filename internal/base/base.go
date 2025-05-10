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

func (s *CombinedMathService) RootMeanSquare(firstValue float64, secondValue float64) (float64, error) {
	xPower, err := s.complexService.Power(firstValue, 2)
	if err != nil {
		return 0, err
	}
	yPower, err := s.complexService.Power(secondValue, 2)
	if err != nil {
		return 0, err
	}
	sum, err := s.simpleService.Add(xPower, yPower)
	if err != nil {
		return 0, err
	}
	avg, err := s.simpleService.Div(sum, 2)
	if err != nil {
		return 0, err
	}
	result, err := s.complexService.Sqrt(avg)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CombinedMathService) GeometricMean(firstValue float64, secondValue float64) (float64, error) {
	product, err := s.simpleService.Mul(firstValue, secondValue)
	if err != nil {
		return 0, err
	}
	result, err := s.complexService.Sqrt(product)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CombinedMathService) BodyMassIndex(weight float64, height float64) (float64, error) {
	heightPower, err := s.complexService.Power(height, 2)
	if err != nil {
		return 0, err
	}
	result, err := s.simpleService.Div(weight, heightPower)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CombinedMathService) PowerLevelDiff(firstValue float64, secondValue float64) (float64, error) {
	ratio, err := s.simpleService.Div(firstValue, secondValue)
	if err != nil {
		return 0, err
	}
	logValue, err := s.complexService.Log(ratio, 10)
	if err != nil {
		return 0, err
	}
	result, err := s.simpleService.Mul(10, logValue)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CombinedMathService) PercentageValueChange(firstValue float64, secondValue float64) (float64, error) {
	delta, err := s.simpleService.Sub(secondValue, firstValue)
	if err != nil {
		return 0, err
	}
	abs, err := s.complexService.Abs(firstValue)
	if err != nil {
		return 0, err
	}
	div, err := s.simpleService.Div(delta, abs)
	if err != nil {
		return 0, err
	}
	percentage, err := s.simpleService.Mul(div, 100)
	if err != nil {
		return 0, err
	}
	result, err := s.complexService.Round(percentage, 2)
	if err != nil {
		return 0, err
	}
	return result, nil
}
