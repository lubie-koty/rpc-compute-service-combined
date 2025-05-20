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

func (s *CombinedMathService) RootMeanSquare(numbers []float64) (float64, error) {
	for idx, num := range numbers {
		power, err := s.complexService.Power(num, 2)
		if err != nil {
			return 0, err
		}
		numbers[idx] = power
	}
	sum, err := s.simpleService.Add(numbers)
	if err != nil {
		return 0, err
	}
	avg, err := s.simpleService.Div([]float64{sum, 2})
	if err != nil {
		return 0, err
	}
	result, err := s.complexService.Sqrt(avg)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CombinedMathService) GeometricMean(numbers []float64) (float64, error) {
	product, err := s.simpleService.Mul(numbers)
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
	result, err := s.simpleService.Div([]float64{weight, heightPower})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CombinedMathService) PowerLevelDiff(firstValue float64, secondValue float64) (float64, error) {
	ratio, err := s.simpleService.Div([]float64{firstValue, secondValue})
	if err != nil {
		return 0, err
	}
	logValue, err := s.complexService.Log(ratio, 10)
	if err != nil {
		return 0, err
	}
	result, err := s.simpleService.Mul([]float64{10, logValue})
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *CombinedMathService) PercentageValueChange(firstValue float64, secondValue float64) (float64, error) {
	delta, err := s.simpleService.Sub([]float64{secondValue, firstValue})
	if err != nil {
		return 0, err
	}
	abs, err := s.complexService.Abs(firstValue)
	if err != nil {
		return 0, err
	}
	div, err := s.simpleService.Div([]float64{delta, abs})
	if err != nil {
		return 0, err
	}
	percentage, err := s.simpleService.Mul([]float64{div, 100})
	if err != nil {
		return 0, err
	}
	result, err := s.complexService.Round(percentage, 2)
	if err != nil {
		return 0, err
	}
	return result, nil
}
