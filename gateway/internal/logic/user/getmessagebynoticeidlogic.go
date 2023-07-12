package user

import (
	"context"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageByNoticeIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageByNoticeIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageByNoticeIdLogic {
	return &GetMessageByNoticeIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetMessageByNoticeId 获取消息
func (l *GetMessageByNoticeIdLogic) GetMessageByNoticeId(req *types.GetMessageByNoticeId) (resp *types.NoticeList, err error) {

	if req.NoticeType == 1 {
		// 获取系统消息
		sysData, err := l.svcCtx.UserRpcClient.QuerySystemNotification(l.ctx, &userclient.QuerySystemNotificationRequest{
			SystemNoticeId: req.NoticeId,
		})
		if err != nil {
			return nil, err
		}
		resp = &types.NoticeList{
			NoticeId:      sysData.SystemNoticeId,
			NoticeType:    1,
			NoticeTitle:   sysData.NoticeTitle,
			NoticeContent: sysData.NoticeContent,
			NoticeTime:    sysData.NoticeStartTime,
		}
		return resp, nil
	} else {
		// 获取用户消息
		usreData, err := l.svcCtx.UserRpcClient.GetUserNotifications(l.ctx, &userclient.GetUserNotificationsRequest{
			UserNoticeId: req.NoticeId,
		})
		if err != nil {
			return nil, err
		}
		resp = &types.NoticeList{
			NoticeId:      usreData.UserId,
			NoticeType:    2,
			NoticeTitle:   "",
			NoticeContent: usreData.NoticeContent,
			NoticeTime:    0,
		}
		return resp, nil
	}
}
