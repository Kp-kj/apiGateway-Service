package curatorial

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLabelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLabelLogic {
	return &DeleteLabelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteLabel 删除标签
func (l *DeleteLabelLogic) DeleteLabel(req *types.TaskIDInquireInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.DeleteLabel(l.ctx, &taskclient.TaskIDInquireInput{Id: req.Id})

	return &types.Mistake{Msg: err1.Msg}, nil
}
