// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"
	"rpc-common/types/user"
	"user/internal/logic"
	"user/internal/svc"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdRequest) (*user.UserResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServer) SaveUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	l := logic.NewUserLogic(ctx, s.svcCtx)
	return l.SaveUser(in)
}
