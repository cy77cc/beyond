package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cy77cc/beyond/application/user/rpc/internal/svc"
	"github.com/cy77cc/beyond/application/user/rpc/pb/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByMobileLogic) FindByMobile(in *service.FindByMobileRequest) (*service.FindByMobileResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		logx.Errorf("FindByMobile user mobile: %s error: %v", in.Mobile, err)
		return nil, status.New(codes.NotFound, "手机号用户不存在").Err()
	}

	return &service.FindByMobileResponse{
		UserId:   user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
	}, nil
}
