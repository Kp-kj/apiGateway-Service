package adminmall

import (
	"context"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCryptominerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCryptominerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCryptominerLogic {
	return &CreateCryptominerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCryptominerLogic) CreateCryptominer(req *types.CreateCryptominerInput) (resp *types.IsSuccessReply, err error) {

	result, err := l.svcCtx.BlockClient.CreateCryptominer(l.ctx, &blockclient.CreateCryptominerRequest{
		AdminuserId:         req.AdminUserID,
		CryptominerName:     req.CryptominerName,
		CryptominerPicture:  req.CryptominerPicture,
		CryptominerDescribe: req.CryptominerDescribe,
		CryptominerPrice:    req.CryptominerPrice,
		CryptominerDuration: req.CryptominerDuration,
		CryptominerPower:    req.CryptominerPower,
		PaymentWay:          req.PaymentWay,
	})

	return &types.IsSuccessReply{IsSuccess: result.IsSuccess}, err
}
