package user

import (
	"context"
	"fmt"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/userclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetUserInfo 获取用户信息逻辑
func (l *GetUserInfoLogic) GetUserInfo() (resp *types.UserInfoReply, err error) {

	//userId := l.ctx.Value("userId").(int64)  // 从上下文中获取用户id

	dbUserInfo, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{
		UserId: "1661640027962085376",
	})

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	fmt.Println(dbUserInfo.UserName)

	return &types.UserInfoReply{
		UserName:    dbUserInfo.UserName,    // 用户名 （会改变）
		TwitterName: dbUserInfo.TwitterName, // twitter名 推特唯一id
		UserAvatar:  dbUserInfo.UserAvatar,  // 用户头像
		IsNew:       dbUserInfo.IsNew,       // 是否是新用户
	}, nil

}
