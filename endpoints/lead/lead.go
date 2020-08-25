package lead

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	reqModel "github.com/vupham79/convertlead-be/model/request/lead"
	"github.com/vupham79/convertlead-be/service"
)

func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req = request.(*reqModel.Create)
		)
		res, err := s.LeadService.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req = request.(*reqModel.Find)
		)
		res, err := s.LeadService.Find(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req = request.(*reqModel.FindAll)
		)
		res, err := s.LeadService.FindAll(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req = request.(*reqModel.Update)
		)
		res, err := s.LeadService.Update(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req = request.(*reqModel.Delete)
		)
		res, err := s.LeadService.Delete(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
