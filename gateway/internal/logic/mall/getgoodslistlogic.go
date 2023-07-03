package mall

import (
	"context"
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

	result, err := l.svcCtx.BlockClient.GetGoodsList(l.ctx, &blockclient.GetGoodsListRequest{
		UserId: req.UserID,
	})

	var cryptominerList []*types.Cryptominer
	for _, item := range result.Cryptominer {
		cryptominerList = append(cryptominerList, &types.Cryptominer{
			CryptominerID:       item.CryptominerId,
			CryptominerTypeID:   item.CryptominerTypeid,
			CryptominerName:     item.CryptominerName,
			CryptominerPicture:  item.CryptominerPicture,
			CryptominerDescribe: item.CryptominerDescribe,
			CryptominerPrice:    item.CryptominerPrice,
			CryptominerDuration: item.CryptominerDuration,
			OptionalStatus:      item.OptionalStatus,
			IsBargain:           item.IsBargain,
		})
	}

	var propList []*types.Prop
	for _, item := range result.Prop {
		propList = append(propList, &types.Prop{
			PropID:       item.PropId,
			PropTypeID:   item.PropTypeid,
			PropName:     item.PropName,
			PropPicture:  item.PropPicture,
			PropDescribe: item.PropDescribe,
			PropPrice:    item.PropPrice,
		})
	}

	return &types.GetGoodsListReply{Cryptominer: cryptominerList, Prop: propList}, err
}
