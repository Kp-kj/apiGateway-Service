package user

import (
	"context"
	"fmt"
	"gateway/userclient"

	"gateway/internal/svc"
	"gateway/internal/types"

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

func (l *GetHelpCategoryListLogic) GetHelpCategoryList(req *types.GetHelpCategoryList) (resp *types.HelpCategoryListReply, err error) {

	helpResp, err := l.svcCtx.UserRpcClient.GetHelpCategories(l.ctx, &userclient.GetHelpCategoriesRequest{CategoryStatus: 0})
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	// 创建帮助分类列表
	categoryList := make([]types.CategoryList, len(helpResp.HelpCategories))
	for i, category := range helpResp.HelpCategories {
		// 查找翻译结果
		translationResp, err := l.svcCtx.UserRpcClient.GetHelpCategoryTranslations(l.ctx, &userclient.GetHelpCategoryTranslationsRequest{
			HelpCategoryId: category.HelpCategoryId,
			Language:       req.LanguageCode,
		})
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}
		// 获取翻译结果
		translation := ""
		if translationResp != nil {
			translation = translationResp.CategoryName
		}

		// 创建分类列表项
		categoryList[i] = types.CategoryList{
			CategoryId:   category.HelpCategoryId,
			CategoryName: translation,
		}
	}
	return &types.HelpCategoryListReply{CategoryList: categoryList}, nil
}
