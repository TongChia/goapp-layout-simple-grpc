package server

import (
	"go.uber.org/fx"
)

var ProviderSet = fx.Provide(NewGRPCServer)
