package adminmall

import (
	"context"
	"fmt"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminActivityListLogic {
	return &AdminActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminActivityListLogic) AdminActivityList(req *types.AdminActivityListInput) (resp *types.AdminActivityListReply, err error) {

	result, err := l.svcCtx.BlockClient.AdminActivityList(l.ctx, &blockclient.AdminActivityListRequest{
		UserId: req.AdminUserID,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var activityList []*types.Activity
	for _, item := range result.Activity {
		activityList = append(activityList, &types.Activity{
			ActivityID:       item.ActivityId,
			CryptominerName:  item.CryptominerName,
			UserAmount:       item.UserAmount,
			MinPrice:         fmt.Sprintf("%.2f", item.MinPrice),
			FirstBargainPer:  fmt.Sprintf("%.2f", item.FirstBargainPer),
			FriendBargainPer: fmt.Sprintf("%.2f", item.FriendBargainPer),
			IsActivation:     item.IsActivation,
		})
	}

	return &types.AdminActivityListReply{Activity: activityList}, err
}
