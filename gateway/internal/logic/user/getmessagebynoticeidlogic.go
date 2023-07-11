package user

import (
	"context"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageByNoticeIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageByNoticeIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByNoticeIdLogic {
	return &GetMessageByNoticeIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetMessageByNoticeId 获取消息
func (l *GetMessageByNoticeIdLogic) GetMessageByNoticeId(req *types.GetMessageByNoticeId) (resp *types.NoticeList, err error) {
	// 先查询 	notices id 的是用户id 还是系统id

	return
}
