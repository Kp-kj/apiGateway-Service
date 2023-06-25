package admin

import (
	"context"

	"gateway/。/internal/svc"
	"gateway/。/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchUpdateUserBlackStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchUpdateUserBlackStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchUpdateUserBlackStatusLogic {
	return &BatchUpdateUserBlackStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchUpdateUserBlackStatusLogic) BatchUpdateUserBlackStatus(req *types.BatchUpdateUserBlackStatus) (resp *types.BatchUpdateUserBlackStatusReply, err error) {
	// todo: add your logic here and delete this line

	return
}
