package admin

import (
	"context"
	"fmt"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserInfoLogic {
	return &AdminUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AdminUserInfo 获取用户信息逻辑
func (l *AdminUserInfoLogic) AdminUserInfo() (resp *types.UserInfoReply, err error) {
	// todo: add your logic here and delete this line
	dbUserInfo, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{
		UserId: "1674266735504527360",
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
