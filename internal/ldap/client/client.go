package client

import (
	"context"
	"time"

	"github.com/panjf2000/gnet/v2"

	"github.com/takumin/gldap/internal/config"
)

type Client struct {
	gnet.BuiltinEventEngine
	context context.Context
	config  *config.Config
}

func NewClient(ctx context.Context, cfg *config.Config) (*Client, error) {
	return &Client{
		config:  cfg,
		context: ctx,
	}, nil
}

func (s *Client) Serve() error {
	return gnet.Run(s, s.config.Client.Endpoint,
		gnet.WithMulticore(true),
		gnet.WithReuseAddr(true),
		gnet.WithReusePort(true),
		gnet.WithTCPKeepAlive(3*time.Second),
		gnet.WithTCPNoDelay(gnet.TCPNoDelay),
	)
}
