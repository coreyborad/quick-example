package handler

import (
	"context"
	"fmt"
	"go-grpc/proto"
	"go-grpc/services"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// DeviceHandler Device Handler
type ExampleHandler struct {
	proto.UnimplementedBasicServiceServer

	serv *services.ExampleService
}

// NewDeviceHandler New Device Handler
func NewExampleHandler(service *services.ExampleService) *ExampleHandler {
	return &ExampleHandler{
		serv: service,
	}
}

// GetTimeAndName GetTimeAndName
func (h *ExampleHandler) GetTimeAndName(ctx context.Context, in *proto.ServerInfo) (*proto.ServerInfo, error) {
	fmt.Println("RECEIVE", in.GetTime(), in.GetName())
	return &proto.ServerInfo{
		Name: "SERVER",
		Time: timestamppb.Now(),
	}, nil
}
