package curatorial

import (
	"context"
	"gateway/taskclient"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ParticipatingTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewParticipatingTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParticipatingTaskLogic {
	return &ParticipatingTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ParticipatingTask 参与策展任务
func (l *ParticipatingTaskLogic) ParticipatingTask(req *types.TaskDetailsInput) (resp *types.Mistake, err error) {
	user, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{UserId: req.UserId})
	if user == nil {
		return nil, err
	}
	err1, err := l.svcCtx.TaskClient.ParticipatingTask(l.ctx, &taskclient.ParticipatingTaskInput{
		UserId:   req.UserId,
		UserName: user.UserName,
		NickName: user.TwitterName,
		Avatar:   user.UserAvatar,
		TaskID:   req.TaskId,
	})
	if err != nil {
		return nil, err
	}
	return &types.Mistake{Msg: err1.Msg}, nil
}
