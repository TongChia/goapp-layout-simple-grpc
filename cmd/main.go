package main

import (
	"context"
	"flag"
	"goapp-layout-template/internal/biz"
	"goapp-layout-template/internal/conf"
	"goapp-layout-template/internal/data"
	"goapp-layout-template/internal/server"
	"goapp-layout-template/internal/service"
	"net"
	"os"

	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs", "config path, eg: -conf ./configs")
}

func newApp(
	conf *conf.Config,
	log grpclog.LoggerV2,
) *fx.App {
	return fx.New(
		fx.Supply(conf),
		fx.Supply(fx.Annotate(log, fx.As(new(grpclog.LoggerV2)))),

		biz.ProviderSet,
		data.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,

		fx.Invoke(run),
	)
}

func run(lifecycle fx.Lifecycle, conf *conf.Config, server *grpc.Server) {
	lifecycle.Append(fx.Hook{OnStart: func(ctx context.Context) error {
		go func() {
			listen, err := net.Listen("tcp", conf.Server.GRPC.Addr)
			if err != nil {
				panic(err)
			}
			if err := server.Serve(listen); err != nil {
				panic(err)
			}
		}()
		return nil
	}, OnStop: func(ctx context.Context) error {
		server.GracefulStop()
		return nil
	}})
}

func main() {
	flag.Parse()

	// parse config
	c, err := conf.ParseConfig(flagconf)
	if err != nil {
		panic(err)
	}

	// init logger
	log := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stderr)

	// create App
	app := newApp(c, log)

	// run App
	app.Run()
}
