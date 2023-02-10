package data

import (
	"context"
	"goapp-layout-template/internal/conf"
	"goapp-layout-template/internal/data/ent"

	"go.uber.org/fx"
	"google.golang.org/grpc/grpclog"
)

// ProviderSet is data providers.
var ProviderSet = fx.Provide(NewData, NewEntClient, NewGreeterRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Config, client *ent.Client, log grpclog.LoggerV2) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
		if err := client.Close(); err != nil {
			log.Error(err)
		}
	}
	return &Data{db: client}, cleanup, nil
}

func NewEntClient(conf *conf.Data, log grpclog.LoggerV2) *ent.Client {
	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
