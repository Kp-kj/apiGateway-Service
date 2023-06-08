package user

import (
	"context"
	"gateway/internal/svc"
	"gateway/internal/types"
	"gateway/userclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelpCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHelpCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelpCategoryListLogic {
	return &GetHelpCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetHelpCategoryList 获取帮助分类列表
func (l *GetHelpCategoryListLogic) GetHelpCategoryList(req *types.GetHelpCategoryList) (resp *types.GetHelpCategoryListReply, err error) {

	// 先获取上架的帮助分类列表   需要改一下
	dbUserInfo, err := l.svcCtx.UserRpcClient.GetHelpCategories(l.ctx, &userclient.GetHelpCategoriesRequest{
		CategoryStatus: 0,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	print(dbUserInfo)
	//根据 languageCode 来请求相应的帮助分类列表

	return nil, nil
}
