package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"rpc-common/types/user"
	"time"

	"userapi/internal/svc"
	"userapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// 超时上下文
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	userResponse, err := l.svcCtx.UserRpc.SaveUser(ctx, &user.UserRequest{
		Name:   req.Name,
		Gender: req.Gender,
	})
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: "success",
		Data:    userResponse,
	}, nil
}

func (l *UserLogic) GerUser(req *types.IdRequest) (resp *types.Response, err error) {
	//认证通过后，从token中获取userId
	userId := l.ctx.Value("userId")
	logx.Infof("获取到的token内容 %s ", userId)
	// 超时上下文
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	userResponse, err := l.svcCtx.UserRpc.GetUser(ctx, &user.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.Response{
		Message: "success",
		Data:    userResponse,
	}
	return resp, nil

}

func (l *UserLogic) getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *UserLogic) Login(req *types.LoginRequest) (resp *types.Response, err error) {
	logx.Info("正在执行login方法")
	userId := 100
	auth := l.svcCtx.Config.Auth
	token, err := l.getToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, int64(userId))
	if err != nil {
		return nil, err
	}
	return &types.Response{
		Message: "success",
		Data:    token,
	}, nil

}
