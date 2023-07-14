package adminmall

import (
	"context"
	"gateway/blockclient"
	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGoodListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGoodListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGoodListLogic {
	return &AdminGoodListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGoodListLogic) AdminGoodList(req *types.AdminGoodListInput) (resp *types.AdminGoodListReply, err error) {

	result, err := l.svcCtx.BlockClient.AdminGoodList(l.ctx, &blockclient.AdminGoodListRequest{
		UserId: req.AdminUserID,
	})
	var adminGoodList []*types.AdminGood
	for _, item := range result.AdminGood {
		adminGoodList = append(adminGoodList, &types.AdminGood{
			GoodTypeID:   item.GoodTypeid,
			GoodName:     item.GoodName,
			GoodDuration: item.GoodDuration,
			PaymentWay:   item.PaymentWay,
			PropPrice:    item.PropPrice,
			GoodStatus:   item.GoodStatus,
			GoodPower:    item.GoodPower,
			GoodType:     item.GoodType,
		})
	}

	return &types.AdminGoodListReply{AdminGood: adminGoodList}, err
}
