package mall

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBargainRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainRecordLogic {
	return &GetBargainRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBargainRecordLogic) GetBargainRecord(req *types.GetBargainRecordInput) (resp *types.GetBargainRecordReply, err error) {
	// todo: add your logic here and delete this line

	return
}
