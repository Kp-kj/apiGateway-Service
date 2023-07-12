package user

import (
	"context"
	"encoding/json"
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
	userId := l.ctx.Value("userId")

	var userIdNum int64
	if v, ok := userId.(json.Number); ok {
		userIdNum, _ = v.Int64()
	}

	//如果req.LastNoticeI为0，则表示是第一次请求，需要从头开始查询
	//如果有值需要转换为string
	//重新赋值一个新变量
	var lastBiticestring string
	if req.LastNoticeId == 0 {
		lastBiticestring = ""
	} else {
		lastBiticestring = fmt.Sprintf("%d", req.LastNoticeId)
	}

	fmt.Println("userId:", userIdNum)
	messageListResp, err := l.svcCtx.UserRpcClient.RecordNotice(l.ctx, &userclient.RecordNoticeRequest{
		UserId:               userIdNum,
		Limit:                30,
		RecordNoticeStatus:   2, //0未读 1已读 2全部
		RecordNoticeCategory: req.NoticeType,
		LastNoticeId:         lastBiticestring,
	})

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	fmt.Println(messageListResp)
	// 构造返回值
	resp = &types.GetMessageListReply{
		NoticeList: make([]types.NoticeList, 0),
	}

	// 将messageListResp中的每个元素赋值到NoticeList中
	for _, notice := range messageListResp.Notice {
		var Notice int64
		if notice.RecordNoticeCategory == 0 {
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
				NoticeType:    1, // 1表示系统通知
				NoticeStatus:  noticeData.NoticeStatus,
			}
			resp.NoticeList = append(resp.NoticeList, noticeItem)
		} else {
			// 如果是用户通知
			Notice = notice.UserNoticeId
			noticeData, err := l.svcCtx.UserRpcClient.GetUserNotifications(l.ctx, &userclient.GetUserNotificationsRequest{
				UserNoticeId: Notice,
			})
			if err != nil {
				return nil, err
			}
			noticeItem := types.NoticeList{
				NoticeId:      Notice, // 消息id
				NoticeTitle:   "",
				NoticeContent: noticeData.NoticeContent,
				NoticeTime:    0, //用户通知没有时间
				NoticeType:    2, // 2表示用户通知
				NoticeStatus:  noticeData.NoticeStatus,
			}
			resp.NoticeList = append(resp.NoticeList, noticeItem)
		}
	}

	return resp, nil
}
