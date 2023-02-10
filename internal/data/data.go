package data

import (
	"goapp-layout-template/internal/conf"

	"go.uber.org/fx"
	"google.golang.org/grpc/grpclog"
)

// ProviderSet is data providers.
var ProviderSet = fx.Provide(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Config, log grpclog.LoggerV2) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
