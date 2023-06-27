package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryChestCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryChestCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryChestCollectionLogic {
	return &QueryChestCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryChestCollectionLogic) QueryChestCollection(req *types.UserIDInquireInput) (resp *types.ReChestCollectionSrt, err error) {
	// todo: add your logic here and delete this line

	return
}
