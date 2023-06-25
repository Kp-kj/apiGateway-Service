package user

import (
	"context"
	"fmt"
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
func (l *GetHelpCategoryListLogic) GetHelpCategoryList() (resp *types.HelpCategoryListReply, err error) {
	// todo: add your logic here and delete this line
	helpResp, err := l.svcCtx.UserRpcClient.GetHelpCategories(l.ctx, &userclient.GetHelpCategoriesRequest{CategoryStatus: 0})

	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	fmt.Println(helpResp.HelpCategories)

	categoryList := make([]types.CategoryList, len(helpResp.HelpCategories))
	for i := range helpResp.HelpCategories {
		categoryList[i] = types.CategoryList{
			CategoryId:   helpResp.HelpCategories[i].HelpCategoryId,
			CategoryName: "helpResp.HelpCategories[i].CategoryName",
		}
	}

	return &types.HelpCategoryListReply{CategoryList: categoryList}, nil
}
