// Code generated by Kitex v0.7.2. DO NOT EDIT.

package signservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	sign "sign/kitex_gen/sign/sign"
)

func serviceInfo() *kitex.ServiceInfo {
	return signServiceServiceInfo
}

var signServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SignService"
	handlerType := (*sign.SignService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Sign": kitex.NewMethodInfo(signHandler, newSignServiceSignArgs, newSignServiceSignResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "sign",
		"ServiceFilePath": `idl/sign.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func signHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceSignArgs)
	realResult := result.(*sign.SignServiceSignResult)
	success, err := handler.(sign.SignService).Sign(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceSignArgs() interface{} {
	return sign.NewSignServiceSignArgs()
}

func newSignServiceSignResult() interface{} {
	return sign.NewSignServiceSignResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Sign(ctx context.Context, req *sign.Empty) (r *sign.Empty, err error) {
	var _args sign.SignServiceSignArgs
	_args.Req = req
	var _result sign.SignServiceSignResult
	if err = p.c.Call(ctx, "Sign", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}