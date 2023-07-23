package stddevapi

import (
	"context"
	"math"
)

type StdDevGroupResult struct {
	Results   []StdDevResult
	ResultSum StdDevResult
}

type StdDevResult struct {
	StdDev float64
	Data   []int
}

func (s *Service) CalculateStdDev(ctx context.Context, requests, length int) (StdDevGroupResult, error) {
	// validate command
	if requests < 1 {
		return StdDevGroupResult{}, NewValidationError(nil, "number of requests must be greater than 0, got %d", requests)
	}
	if length < 1 {
		return StdDevGroupResult{}, NewValidationError(nil, "length must be greater than 0, got %d", length)
	}

	dataGroup, err := s.getRandomIntegersAsync(ctx, requests, length)
	if err != nil {
		return StdDevGroupResult{}, err
	}

	return getResultFromDataGroup(dataGroup), nil
}

func getResultFromDataGroup(dataGroup [][]int) StdDevGroupResult {
	res := StdDevGroupResult{
		Results: make([]StdDevResult, 0),
	}

	for _, data := range dataGroup {
		res.Results = append(res.Results, StdDevResult{
			Data:   data,
			StdDev: calcStdDev(data),
		})

		res.ResultSum.Data = append(res.ResultSum.Data, data...)
	}

	res.ResultSum.StdDev = calcStdDev(res.ResultSum.Data)

	return res
}

func calcStdDev(data []int) (sd float64) {
	mean := calcMean(data)

	for _, d := range data {
		sd += math.Pow(float64(d)-mean, 2)
	}

	return math.Sqrt(sd / float64(len(data)))
}

func calcMean(data []int) float64 {
	var sum float64
	for _, d := range data {
		sum += float64(d)
	}

	return sum / float64(len(data))
}
