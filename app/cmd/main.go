package main

import (
	"github.com/mf-sakura/bh_user/app/config"
	"github.com/mf-sakura/bh_user/app/db"
	upb "github.com/mf-sakura/bh_user/app/proto"
	"github.com/mf-sakura/bh_user/app/server"

	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"net"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	open_config "github.com/uber/jaeger-client-go/config"
	"time"
)

const (
	port = ":5002"
)

func main() {
	fmt.Println("Process Started.")
	cfg := open_config.Configuration{
		Sampler: &open_config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &open_config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "jaeger:6831",
		},
	}
	tracer, closer, err := cfg.New(
		"booking_hotel",
		open_config.Logger(jaeger.StdLogger),
	)
	defer closer.Close()
	if err != nil {
		fmt.Println(err)
	}
	opentracing.SetGlobalTracer(tracer)
	conf, err := config.LoadConifg()
	if err != nil {
		panic(err)
	}
	dsn, err := db.CreateDataSourceName(conf.Port, conf.Host, "bh_user", conf.User, conf.Password)
	if err != nil {
		panic(err)
	}
	if err := db.NewDB(dsn); err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_opentracing.UnaryServerInterceptor()))

	defer func() {
		err := recover()
		s.GracefulStop()
		if err != nil {
			panic(err)
		}
	}()
	upb.RegisterUserServiceServer(s, &server.UserServiceServerImpl{})

	if err := s.Serve(listen); err != nil {
		panic(err)
	}
}
