package middlewares

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (*UserMiddleware) LoginAndReg(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("调用login和register之前执行")
		next(w, r)
		logx.Info("调用login和register之后执行")
	}
}

func (*UserMiddleware) Global(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("all 之前执行")
		next(w, r)
		logx.Info("all 之后执行")
	}
}
