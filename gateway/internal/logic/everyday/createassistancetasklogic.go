package everyday

import (
	"context"
	"gateway/taskclient/task"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAssistanceTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAssistanceTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAssistanceTaskLogic {
	return &CreateAssistanceTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAssistanceTaskLogic) CreateAssistanceTask(req *types.UserIDInquireInput) (resp *types.Mistake, err error) {
	// 获取用户信息
	user, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	// 传输数据进行创建
	err1, err := l.svcCtx.TaskClient.CreateAssistanceTask(l.ctx, &task.CreateUserPublishingAssistanceTaskInput{
		UserId:   req.UserId,
		UserName: user.UserName,
		Avatar:   user.UserAvatar,
	})

	return &types.Mistake{Msg: err1.Msg}, err
}
