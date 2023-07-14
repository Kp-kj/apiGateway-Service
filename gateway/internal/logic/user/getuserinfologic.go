package user

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/userclient"
	"strconv"

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
	userId := l.ctx.Value("userId")
	var userIdStr string
	if v, ok := userId.(json.Number); ok {
		userIdNum, _ := v.Int64()
		userIdStr = strconv.FormatInt(userIdNum, 10)
	}

	dbUserInfo, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{
		UserId: userIdStr,
	})

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	fmt.Println(dbUserInfo.UserName)

	return &types.UserInfoReply{
		UserId:      userIdStr,              // 用户id
		UserName:    dbUserInfo.UserName,    // 用户名 （会改变）
		TwitterName: dbUserInfo.TwitterName, // twitter名 推特唯一id
		UserAvatar:  dbUserInfo.UserAvatar,  // 用户头像
		IsNew:       dbUserInfo.IsNew,       // 是否是新用户
	}, nil

}
