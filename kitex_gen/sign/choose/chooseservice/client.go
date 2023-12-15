// Code generated by Kitex v0.7.2. DO NOT EDIT.

package chooseservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	choose "sign/kitex_gen/sign/choose"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Choose(ctx context.Context, req *choose.Empty, callOptions ...callopt.Option) (r *choose.Empty, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kChooseServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kChooseServiceClient struct {
	*kClient
}

func (p *kChooseServiceClient) Choose(ctx context.Context, req *choose.Empty, callOptions ...callopt.Option) (r *choose.Empty, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Choose(ctx, req)
}
