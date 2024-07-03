package handler

import (
	"userapi/internal/svc"
)

type UserHandler struct {
	serCtx *svc.ServiceContext
}

func NewUserHandler(serCtx *svc.ServiceContext) *UserHandler {
	return &UserHandler{
		serCtx: serCtx,
	}
}
