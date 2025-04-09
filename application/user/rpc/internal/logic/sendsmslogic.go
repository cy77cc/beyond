package logic

import (
	"context"

	"github.com/cy77cc/beyond/application/user/rpc/internal/svc"
	"github.com/cy77cc/beyond/application/user/rpc/pb/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *service.SendSmsRequest) (*service.SendSmsResponse, error) {
	// todo: send sms

	return &service.SendSmsResponse{}, nil
}
