package curatorial

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLabelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLabelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLabelListLogic {
	return &QueryLabelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLabelListLogic) QueryLabelList(req *types.UserIDInquireInput) (resp *types.ReLabelListOut, err error) {
	label, err := l.svcCtx.TaskClient.QueryLabelList(l.ctx, &taskclient.UserIDInquireInput{UserId: req.UserId})
	var labelList []*types.ReLabelList
	for _, item := range label.ReLabelListOut {
		labelList = append(labelList, &types.ReLabelList{
			Id:      item.Id,
			Creator: item.Creator,
			Content: item.Content,
		})

	}
	return &types.ReLabelListOut{ReLabelList: labelList}, nil
}
