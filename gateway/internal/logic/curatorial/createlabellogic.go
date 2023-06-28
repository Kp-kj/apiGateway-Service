package curatorial

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLabelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLabelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLabelLogic {
	return &CreateLabelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateLabel 创建标签
func (l *CreateLabelLogic) CreateLabel(req *types.CreateLabelInput) (resp *types.Mistake, err error) {
	err1, err := l.svcCtx.TaskClient.CreateLabel(l.ctx, &taskclient.CreateLabelInput{
		UserId: req.UserId,
		Label:  req.Label,
	})

	return &types.Mistake{Msg: err1.Msg}, err
}
