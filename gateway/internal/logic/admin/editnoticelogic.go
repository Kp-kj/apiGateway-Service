package admin

import (
	"context"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/userclient/user"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditNoticeLogic {
	return &EditNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 时间转换
func convertTimeToUnix(timestampStr string) (int64, error) {
	if timestampStr == "" {
		return 0, nil // 返回0表示空时间戳，不视为空字符串的解析错误
	}
	t, err := time.Parse("2006-01-02 15:04:05", timestampStr)
	if err != nil {
		return 0, err
	}
	unixTimestamp := t.Unix()
	return unixTimestamp, nil
}

// 编辑公告
func (l *EditNoticeLogic) EditNotice(req *types.EditNotice) (resp *types.EditNoticeReply, err error) {
	var Category int64
	if req.IsAutoPublish {
		Category = 0
	} else {
		Category = 1
	}
	NoticeStartTime, err := convertTimeToUnix(req.SystemNoticeTime)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if req.SystemNoticeId == 0 {
		//如果不传公告id则创建公告
		_, err = l.svcCtx.UserRpcClient.CreateSystemNotification(l.ctx, &user.CreateSystemNotificationRequest{
			NoticeStatus:    req.SystemNoticeStatus,  //公告状态 0-上架 1-下架
			NoticeTitle:     req.SystemNoticeTitle,   //公告标题
			NoticeContent:   req.SystemNoticeContent, //公告内容
			NoticeCategory:  Category,                //公告类别 0-自动发布 1-手动发布
			NoticeStartTime: NoticeStartTime,         //公告开始时间
		})
		if err != nil {
			logx.Error(err)
			return nil, err
		}

	} else {
		//如果传了公告id则更新公告
		_, err = l.svcCtx.UserRpcClient.EditSystemNotification(l.ctx, &user.EditSystemNotificationRequest{
			SystemNoticeId:  req.SystemNoticeId,      //公告id
			NoticeStatus:    req.SystemNoticeStatus,  //公告状态 0-上架 1-下架
			NoticeTitle:     req.SystemNoticeTitle,   //公告标题
			NoticeContent:   req.SystemNoticeContent, //公告内容
			NoticeCategory:  Category,                //公告类别 0-自动发布 1-手动发布
			NoticeStartTime: NoticeStartTime,         //公告开始时间
		})
	}
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &types.EditNoticeReply{
		IsSuccess: true,
	}, nil
}
