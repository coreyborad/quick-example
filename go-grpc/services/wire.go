// +build wireinject

package services

import (
	"github.com/google/wire"
)

var (
	ExampleServiceSet = wire.NewSet(NewExampleService)
)

// CreateExampleService CreateExampleService
func CreateExampleService() *ExampleService {
	wire.Build(ExampleServiceSet)

	return nil
}
