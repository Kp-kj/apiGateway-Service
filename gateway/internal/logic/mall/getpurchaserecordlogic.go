package mall

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/blockclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPurchaseRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPurchaseRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPurchaseRecordLogic {
	return &GetPurchaseRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPurchaseRecordLogic) GetPurchaseRecord(req *types.GetPurchaseRecordInput) (resp *types.GetPurchaseRecordReply, err error) {
	userId := l.ctx.Value("userId")
	var userID int64
	if v, ok := userId.(json.Number); ok {
		userID, _ = v.Int64()
	}
	record, err := l.svcCtx.BlockClient.GetPurchaseRecord(l.ctx, &blockclient.GetPurchaseRecordRequest{
		UserId: userID,
	})
	var PurchaseRecordList []*types.PurchaseRecord
	for _, item := range record.PurchaseRecord {
		PurchaseRecordList = append(PurchaseRecordList, &types.PurchaseRecord{
			PurchaseRecordID: item.PurchaseRecordId,
			GoodName:         item.GoodName,
			GoodPicture:      item.GoodPicture,
			GoodQuantity:     item.GoodQuantity,
			PurchaseTime:     item.PurchaseTime,
			PurchasePrice:    fmt.Sprintf("%.2f", item.PurchasePrice),
			PurchaseWay:      item.PurchaseWay,
			PaymentWay:       item.PaymentWay,
		})
	}

	return &types.GetPurchaseRecordReply{PurchaseRecord: PurchaseRecordList}, err
}
