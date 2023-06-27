package everyday

import (
	"context"

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

func (l *AmendChestCollectionLogic) AmendChestCollection(req *types.AmendChestCollectionInput) (resp *types.Mistake, err error) {
	// todo: add your logic here and delete this line

	return
}
