package admin

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

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

// BatchUpdateUserBlackStatus 批量更新用户黑名单状态
func (l *BatchUpdateUserBlackStatusLogic) BatchUpdateUserBlackStatus(req *types.BatchUpdateUserBlackStatus) (resp *types.BatchUpdateUserBlackStatusReply, err error) {
	// todo: add your logic here and delete this line

	return &types.BatchUpdateUserBlackStatusReply{
		UserIdList:  []int64{1, 2, 3},
		BlackStatus: 0,
	}, nil
}
