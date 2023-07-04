package mall

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainProgressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBargainProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainProgressLogic {
	return &GetBargainProgressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBargainProgressLogic) GetBargainProgress(req *types.GetBargainProgressInput) (resp *types.GetBargainProgressReply, err error) {
	// todo: add your logic here and delete this line

	return
}
