package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoint : single Remote Procedure Call. We write an adapter for each method

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(ctx, req.S)

		if err != nil {
			return uppercaseResponse{V: v, Err: err.Error()}, nil
		}

		return uppercaseResponse{V: v, Err: ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(ctx, req.S)

		return countResponse{V: v}, nil
	}
}
