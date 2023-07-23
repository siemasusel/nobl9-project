package stddevapi

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type RandomBackend interface {
	GetRandomIntegers(ctx context.Context, length int) ([]int, error)
}

type Service struct {
	randBackend RandomBackend
}

func NewService(intBackend RandomBackend) *Service {
	return &Service{
		randBackend: intBackend,
	}
}

func (s *Service) getRandomIntegersAsync(ctx context.Context, requests, length int) ([][]int, error) {
	dataCh := make(chan []int, requests)
	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < requests; i++ {
		g.Go(func() error {
			data, err := s.randBackend.GetRandomIntegers(ctx, length)
			if err != nil {
				return err
			}

			dataCh <- data

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	close(dataCh)

	dataGroup := make([][]int, 0, requests)
	for data := range dataCh {
		dataGroup = append(dataGroup, data)
	}

	return dataGroup, nil
}
