package everyday

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AmendChestCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAmendChestCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AmendChestCollectionLogic {
	return &AmendChestCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AmendChestCollection 创建宝箱领取度+更新宝箱领取进度
func (l *AmendChestCollectionLogic) AmendChestCollection(req *types.AmendChestCollectionInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.AmendChestCollection(l.ctx, &taskclient.AmendChestCollectionInput{
		UserId: req.UserId,
		Amount: req.Amount,
	})
	return &types.Mistake{Msg: err1.Msg}, err
}
