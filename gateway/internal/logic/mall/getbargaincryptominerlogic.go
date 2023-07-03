package mall

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainCryptominerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBargainCryptominerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainCryptominerLogic {
	return &GetBargainCryptominerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBargainCryptominerLogic) GetBargainCryptominer(req *types.GetBargainCryptominerInput) (resp *types.GetBargainCryptominerReply, err error) {
	// todo: add your logic here and delete this line

	return
}
