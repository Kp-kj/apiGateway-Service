package everyday

import (
	"context"
	"fmt"
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
	fmt.Printf("111111111111111111\n")
	err1, err := l.svcCtx.TaskClient.AmendChestCollection(l.ctx, &taskclient.AmendChestCollectionInput{
		UserId: req.UserId,
		Amount: req.Amount,
	})
	fmt.Printf("22222222222222222\n")
	return &types.Mistake{Msg: err1.Msg}, err
}
