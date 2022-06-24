package server

import (
	"context"
	"time"

	"github.com/panjf2000/gnet/v2"

	"github.com/takumin/gldap/internal/config"
)

type Server struct {
	gnet.BuiltinEventEngine
	context context.Context
	config  *config.Config
}

func NewServer(ctx context.Context, cfg *config.Config) (*Server, error) {
	return &Server{
		config:  cfg,
		context: ctx,
	}, nil
}

func (s *Server) Serve() error {
	return gnet.Run(s, s.config.Server.ListenURL,
		gnet.WithMulticore(true),
		gnet.WithReuseAddr(true),
		gnet.WithReusePort(true),
		gnet.WithTCPKeepAlive(3*time.Second),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
	)
}
