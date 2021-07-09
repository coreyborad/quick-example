// +build wireinject

package handler

import (
	"go-grpc/services"

	"github.com/google/wire"
)

var (
	ExampleHandlerSet = wire.NewSet(NewExampleHandler, services.CreateExampleService)
)

// CreateExampleHandler CreateExampleHandler
func CreateExampleHandler() *ExampleHandler {
	wire.Build(ExampleHandlerSet)

	return nil
}
