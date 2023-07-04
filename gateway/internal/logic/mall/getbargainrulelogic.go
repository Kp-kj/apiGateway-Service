package mall

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainRuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBargainRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainRuleLogic {
	return &GetBargainRuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBargainRuleLogic) GetBargainRule(req *types.GetBargainRuleInput) (resp *types.GetBargainRuleReply, err error) {
	// todo: add your logic here and delete this line

	return
}
