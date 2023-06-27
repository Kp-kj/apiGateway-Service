package everyday

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChestCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChestCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChestCollectionLogic {
	return &GetChestCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChestCollectionLogic) GetChestCollection(req *types.UserIDInquireInput) (resp *types.ReChestCollectionSrt, err error) {
	// todo: add your logic here and delete this line

	return
}
