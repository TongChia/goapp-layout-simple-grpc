package biz

import (
	"go.uber.org/fx"
)

var ProviderSet = fx.Provide(NewGreeterUsecase)
