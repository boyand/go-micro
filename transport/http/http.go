package http

import (
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/transport"
)

func init() {
	cmd.DefaultTransports["http"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewTransport(opts...)
}
