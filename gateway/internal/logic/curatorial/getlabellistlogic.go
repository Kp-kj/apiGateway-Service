package curatorial

import (
	"context"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/taskclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLabelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLabelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLabelListLogic {
	return &GetLabelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetLabelList 获取标签列表
func (l *GetLabelListLogic) GetLabelList(req *types.UserIDInquireInput) (resp *types.ReLabelListOut, err error) {
	// todo: add your logic here and delete this line
	reLabelListOut, err := l.svcCtx.TaskClient.QueryLabelList(l.ctx, &taskclient.UserIDInquireInput{
		UserId: req.UserId})
	var reLabelList []*types.ReLabelList
	for _, item := range reLabelListOut.ReLabelListOut {
		reLabelList = append(reLabelList, &types.ReLabelList{
			Id:      item.Id,
			Creator: item.Creator,
			Content: item.Content,
		})
	}
	return &types.ReLabelListOut{
		ReLabelList: reLabelList,
	}, nil
}