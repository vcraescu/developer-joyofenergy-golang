package readings

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"joi-energy-golang/domain"
)

func makeStoreReadingsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.StoreReadings)
		s.StoreReadings(req.SmartMeterId, req.ElectricityReadings)

		return req, nil
	}
}

func makeGetReadingsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		res := s.GetReadings(req)

		if res == nil {
			return nil, domain.ErrNotFound
		}

		return domain.StoreReadings{
			SmartMeterId:        req,
			ElectricityReadings: res,
		}, nil
	}
}
