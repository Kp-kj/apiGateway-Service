package user

import (
	"context"
	"fmt"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetMessageList 获取消息列表
func (l *GetMessageListLogic) GetMessageList(req *types.GetMessageList) (resp *types.GetMessageListReply, err error) {

	fmt.Println("req: ", req)
	userId := l.ctx.Value("userId").(int64) // 从上下文中获取用户id

	fmt.Println("userId: ", userId)
	messageListResp, err := l.svcCtx.UserRpcClient.RecordNotice(l.ctx, &userclient.RecordNoticeRequest{
		UserId:               userId,
		Limit:                30,
		RecordNoticeStatus:   2, //0未读 1已读 2全部
		RecordNoticeCategory: req.NoticeType,
		LastNoticeId:         string(req.LastNoticeId),
	})

	if err != nil {
		return nil, err
	}
	// 构造返回值
	resp = &types.GetMessageListReply{
		NoticeList: make([]types.NoticeList, 0),
	}

	// 将messageListResp中的每个元素赋值到NoticeList中
	for _, notice := range messageListResp.Notice {
		var Notice int64
		if notice.RecordNoticeCategory == 1 {
			// 如果是系统通知
			Notice = notice.SystemNoticeId
			noticeData, err := l.svcCtx.UserRpcClient.QuerySystemNotification(l.ctx, &userclient.QuerySystemNotificationRequest{
				SystemNoticeId: Notice,
			})
			if err != nil {
				return nil, err
			}
			// 构造NoticeList对象并赋值
			noticeItem := types.NoticeList{
				NoticeId:      Notice, // 消息id
				NoticeTitle:   noticeData.NoticeTitle,
				NoticeContent: noticeData.NoticeContent,
				NoticeTime:    noticeData.NoticeStartTime,
				NoticeType:    noticeData.NoticeCategory,
				NoticeStatus:  noticeData.NoticeStatus,
			}
			resp.NoticeList = append(resp.NoticeList, noticeItem)
		} else {
			// 如果是用户通知
			Notice = notice.UserNoticeId
		}
	}

	return resp, nil
}
