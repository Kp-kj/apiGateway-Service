package admin

import (
	"context"
	"fmt"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditNoticeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditNoticeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditNoticeStatusLogic {
	return &EditNoticeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// EditNoticeStatus 修改公告状态
func (l *EditNoticeStatusLogic) EditNoticeStatus(req *types.EditNoticeStatus) (resp *types.EditNoticeStatusReply, err error) {
	fmt.Println(req.EditItems)

	//l.svcCtx.UserRpcClient.

	return
}
