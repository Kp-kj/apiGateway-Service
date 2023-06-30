package everyday

import (
	"context"
	"gateway/taskclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAssistanceTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAssistanceTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAssistanceTaskLogic {
	return &QueryAssistanceTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryAssistanceTask 查询用户发布助力信息
func (l *QueryAssistanceTaskLogic) QueryAssistanceTask(req *types.UserIDInquireInput) (resp *types.UserPublishingAssistanceTask, err error) {
	data, err := l.svcCtx.TaskClient.QueryAssistanceTask(l.ctx, &taskclient.UserIDInquireInput{UserId: req.UserId})
	return &types.UserPublishingAssistanceTask{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		UserId:    data.UserId,
		UserName:  data.UserName,
		Avatar:    data.Avatar,
		Article:   data.Article,
		Link:      data.Link,
		Label:     data.Label,
	}, nil
}
