package mall

import (
	"context"
	"encoding/json"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/userclient"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBargainRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainRecordLogic {
	return &GetBargainRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBargainRecordLogic) GetBargainRecord(req *types.GetBargainRecordInput) (resp *types.GetBargainRecordReply, err error) {
	userId := l.ctx.Value("userId")
	var userID int64
	var LoginStatus string = "0"
	if v, ok := userId.(json.Number); ok {
		userID, _ = v.Int64()
	}

	user, err := l.svcCtx.UserRpcClient.QueryUser(l.ctx, &userclient.QueryUserRequest{
		UserId: strconv.Itoa(int(userID)),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	record, err := l.svcCtx.BlockClient.GetBargainRecord(l.ctx, &blockclient.GetBargainRecordRequest{
		UserId:    userID,
		BargainId: req.BargainID,
	})
	var SupportUserList []*types.SupportUser
	for _, supportUser := range record.BargainRecord.SupportUser {
		oneSupportUser := &types.SupportUser{
			AssistorID:   supportUser.AssistorId,
			AssistorName: supportUser.AssistorName,
			TwitterName:  supportUser.TwitterName,
			Avatar:       supportUser.Avatar,
			BargainPrice: float64(supportUser.BargainPrice),
		}
		SupportUserList = append(SupportUserList, oneSupportUser)
	}

	BargainRecordList := &types.BargainRecord{
		BargainID:          record.BargainRecord.BargainId,
		UserName:           user.UserName,
		TwitterName:        user.TwitterName,
		CryptominerName:    record.BargainRecord.CryptominerName,
		CryptominerPicture: record.BargainRecord.CryptominerPicture,
		CryptominerPrice:   record.BargainRecord.CryptominerPrice,
		ActivityStartTime:  record.BargainRecord.ActivityStartTime,
		RemainingPrice:     float64(record.BargainRecord.RemainingPrice),
		SupportUser:        SupportUserList,
	}
	LoginStatus = record.LoginStatus
	if userID == 0 {
		LoginStatus = "2"
	}
	return &types.GetBargainRecordReply{BargainRecord: BargainRecordList, LoginStatus: LoginStatus, InActivity: record.InActivity, BargainMax: record.BargainMax, ISBargain: record.IsBargain}, err
}
