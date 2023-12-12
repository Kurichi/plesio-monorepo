//go:build wireinject

package register

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func New() *grpc.Server {
	wire.Build(
		Register,
	)

	return nil
}
