package mall

import (
	"context"
	"encoding/json"
	"gateway/blockclient"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsListLogic) GetGoodsList(req *types.GetGoodsListInput) (resp *types.GetGoodsListReply, err error) {
	userId := l.ctx.Value("userId")
	var userID int64
	if v, ok := userId.(json.Number); ok {
		userID, _ = v.Int64()
	}

	result, err := l.svcCtx.BlockClient.GetGoodsList(l.ctx, &blockclient.GetGoodsListRequest{
		UserId: userID,
	})
	var cryptominerList []*types.Cryptominer
	for _, item := range result.Cryptominer {
		cryptominerList = append(cryptominerList, &types.Cryptominer{
			UserID:              item.UserId,
			CryptominerID:       item.CryptominerId,
			CryptominerTypeID:   item.CryptominerTypeid,
			CryptominerName:     item.CryptominerName,
			CryptominerPicture:  item.CryptominerPicture,
			CryptominerDescribe: item.CryptominerDescribe,
			CryptominerPrice:    item.CryptominerPrice,
			CryptominerDuration: item.CryptominerDuration,
			CryptominerPower:    item.CryptominerPower,
			PaymentWay:          item.PaymentWay,
			OptionalStatus:      item.OptionalStatus,
			IsBargain:           item.IsBargain,
		})
	}

	var propList []*types.Prop
	for _, item := range result.Prop {
		propList = append(propList, &types.Prop{
			UserID:       item.UserId,
			PropID:       item.PropId,
			PropTypeID:   item.PropTypeid,
			PropName:     item.PropName,
			PropPicture:  item.PropPicture,
			PropDescribe: item.PropDescribe,
			PropPrice:    item.PropPrice,
			PaymentWay:   item.PaymentWay,
		})
	}

	return &types.GetGoodsListReply{Cryptominer: cryptominerList, Prop: propList}, err
}
