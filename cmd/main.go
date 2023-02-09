package main

import (
	"flag"
	"os"

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
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger grpclog.LoggerV2, gs *grpc.Server) {
}

func main() {
	flag.Parse()
}
