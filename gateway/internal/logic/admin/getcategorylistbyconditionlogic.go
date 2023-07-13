package admin

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListByConditionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListByConditionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListByConditionLogic {
	return &GetCategoryListByConditionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCategoryListByConditionLogic) GetCategoryListByCondition(req *types.GetCategoryListByCondition) (resp *types.GetCategoryListByConditionReply, err error) {
	// todo: add your logic here and delete this line

	return
}
