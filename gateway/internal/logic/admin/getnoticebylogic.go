package admin

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeByLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoticeByLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeByLogic {
	return &GetNoticeByLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoticeByLogic) GetNoticeBy(req *types.GetNoticeBy) (resp *types.GetNoticeByReply, err error) {
	// todo: add your logic here and delete this line

	return
}
