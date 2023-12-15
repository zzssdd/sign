package main

import (
	"context"
	base "sign/kitex_gen/sign/base"
	choose "sign/kitex_gen/sign/choose"
	sign "sign/kitex_gen/sign/sign"
)

// BaseServiceImpl implements the last service interface defined in the IDL.
type BaseServiceImpl struct{}

// Register implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Register(ctx context.Context, req *base.UserInfo) (resp *base.BaseResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Login(ctx context.Context, req *base.UserInfo) (resp *base.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// Join implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Join(ctx context.Context, req *base.JoinReq) (resp *base.BaseResp, err error) {
	// TODO: Your code here...
	return
}

// CreateGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) CreateGroup(ctx context.Context, req *base.GroupInfo) (resp *base.BaseResp, err error) {
	// TODO: Your code here...
	return
}

// Sign implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Sign(ctx context.Context, req *base.SignReq) (resp *base.BaseResp, err error) {
	// TODO: Your code here...
	return
}

// SignMonth implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) SignMonth(ctx context.Context, req *base.MonthSignReq) (resp *base.MonthSignResp, err error) {
	// TODO: Your code here...
	return
}

// Choose implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Choose(ctx context.Context, req *base.ChooseReq) (resp *base.ChooseResp, err error) {
	// TODO: Your code here...
	return
}

// ChooseSubmit implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChooseSubmit(ctx context.Context, req *base.ChooseSubmitReq) (resp *base.ChooseSubmitResp, err error) {
	// TODO: Your code here...
	return
}

// AddActivity implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) AddActivity(ctx context.Context, req *base.ActicityInfo) (resp *base.BaseResp, err error) {
	// TODO: Your code here...
	return
}

// AddPrize implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) AddPrize(ctx context.Context, req *base.PrizeInfo) (resp *base.BaseResp, err error) {
	// TODO: Your code here...
	return
}

// Choose implements the ChooseServiceImpl interface.
func (s *ChooseServiceImpl) Choose(ctx context.Context, req *choose.Empty) (resp *choose.Empty, err error) {
	// TODO: Your code here...
	return
}

// Sign implements the SignServiceImpl interface.
func (s *SignServiceImpl) Sign(ctx context.Context, req *sign.Empty) (resp *sign.Empty, err error) {
	// TODO: Your code here...
	return
}
