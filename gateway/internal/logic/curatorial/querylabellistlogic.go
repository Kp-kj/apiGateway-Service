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

// QueryLabelList 查询标签列表
func (l *QueryLabelListLogic) QueryLabelList(req *types.UserIDInquireInput) (resp *types.ReLabelListOut, err error) {
	reLabelListOut, err := l.svcCtx.TaskClient.QueryLabelList(l.ctx, &taskclient.UserIDInquireInput{
		UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	var reLabelList []*types.ReLabelList
	for _, item := range reLabelListOut.ReLabelList {
		reLabelList = append(reLabelList, &types.ReLabelList{
			Id:      item.Id,
			Creator: item.Creator,
			Content: item.Content,
		})
	}
	return &types.ReLabelListOut{
		ReLabelList: reLabelList,
	}, err
}
