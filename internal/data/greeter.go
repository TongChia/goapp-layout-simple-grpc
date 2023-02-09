package data

import (
	"context"
	"goapp-layout-template/internal/biz"

	"google.golang.org/grpc/grpclog"
)

type greeterRepo struct {
	data *Data
	log  grpclog.LoggerV2
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, log grpclog.LoggerV2) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log,
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
