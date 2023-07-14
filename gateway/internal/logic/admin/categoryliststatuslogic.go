package admin

import (
	"context"
	"gateway/userclient/user"

	"gateway/internal/svc"
	"gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategorylistStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategorylistStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategorylistStatusLogic {
	return &CategorylistStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CategorylistStatus 批量修改分类状态
func (l *CategorylistStatusLogic) CategorylistStatus(req *types.CategorylistStatus) (resp *types.CategorylistStatusReply, err error) {
	//获取所有传入的分类id
	for _, v := range req.CategoryIds {
		//调用修改分类状态接口
		_, err = l.svcCtx.UserRpcClient.EditHelpCategory(l.ctx, &user.EditHelpCategoryRequest{HelpCategoryId: v, CategoryStatus: req.Status})
		if err != nil {
			return nil, err
		}
	}
	return &types.CategorylistStatusReply{IsSuccess: true}, nil
}
