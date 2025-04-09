package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cy77cc/beyond/application/user/rpc/internal/svc"
	"github.com/cy77cc/beyond/application/user/rpc/pb/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *service.FindByIdRequest) (*service.FindByIdResponse, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)

	if err != nil {
		logx.Errorf("FindById userId: %d error: %v", in.UserId, err)
		return nil, status.New(codes.NotFound, "用户不存在").Err()
	}

	return &service.FindByIdResponse{
		UserId:   user.Id,
		Username: user.Username,
		//Mobile:   user.Mobile,
		Avatar: user.Avatar,
	}, nil
}
