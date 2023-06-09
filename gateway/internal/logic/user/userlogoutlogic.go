package user

import (
	"context"
	"fmt"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogoutLogic {
	return &UserLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogoutLogic) UserLogout() (resp *types.UserLoginReply, err error) {
	// todo: 编写用户退出逻辑  暂时不做任何操作
	fmt.Println(l.ctx.Value("userId")) // 从jwt token中解析出来的userId
	return &types.UserLoginReply{
		Token: "",
	}, nil
}
