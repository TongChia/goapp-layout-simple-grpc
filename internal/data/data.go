package data

import (
	"goapp-layout-template/internal/conf"

	"google.golang.org/grpc/grpclog"
)

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
