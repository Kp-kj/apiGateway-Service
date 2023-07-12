package admin

import (
	"context"
	"fmt"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubtaskStyleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubtaskStyleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubtaskStyleLogic {
	return &QuerySubtaskStyleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySubtaskStyle 查询子任务样式列表
func (l *QuerySubtaskStyleLogic) QuerySubtaskStyle(req *types.TaskIDInquireInput) (resp *types.ReSubtaskStyle, err error) {
	fmt.Printf("!1111111111111111111111\n")
	data, err := l.svcCtx.TaskClient.QuerySubtaskStyle(l.ctx, &taskclient.TaskIDInquireInput{TaskId: req.TaskId})
	if err != nil {
		fmt.Printf("22222222:%v\n", err)
		return nil, err
	}
	var subtaskStyle []*types.SubtaskStyle
	for _, item := range data.SubtaskStyle {
		subtaskStyle = append(subtaskStyle, &types.SubtaskStyle{
			TaskId:         item.TaskId,
			TaskName:       item.TaskName,
			TaskNameEng:    item.TaskNameEng,
			TaskDetails:    item.TaskDetails,
			TaskDetailsEng: item.TaskDetailsEng,
			TaskStatus:     item.TaskStatus,
		})
	}
	return &types.ReSubtaskStyle{SubtaskStyle: subtaskStyle}, err
}
