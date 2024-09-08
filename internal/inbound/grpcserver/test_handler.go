package grpcserver

import (
	"context"
	"fmt"
	"my-app/internal/inbound/grpcserver/proto/pb/testendpoint"
	"my-app/internal/pkg/services"
)

type TestEndpointServer struct {
	testendpoint.UnimplementedTestEndpointServer
}

func (s *TestEndpointServer) GetIdsByName(ctx context.Context, req *testendpoint.GetIdsByNameRequest) (*testendpoint.GetIdsByNameResponse, error) {
	ids, err := services.GetTestIdsByName(req.Name)
	if err != nil {
		return nil, fmt.Errorf("error getting ids for name: %s", req.Name)
	}
	return &testendpoint.GetIdsByNameResponse{Ids: ids}, nil
}
